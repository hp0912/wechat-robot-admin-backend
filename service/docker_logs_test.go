package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"wechat-robot-admin-backend/model"

	"github.com/docker/docker/api/types/filters"
)

func TestRobotContainerLogsKeepAvailableContainerLogs(t *testing.T) {
	for _, tc := range []struct {
		name          string
		clientMissing bool
		serverMissing bool
		clientFailure bool
	}{
		{name: "两端日志均可读取"},
		{name: "客户端不存在仍读取服务端", clientMissing: true},
		{name: "服务端不存在保留客户端", serverMissing: true},
		{name: "客户端读取失败仍读取服务端", clientFailure: true},
		{name: "两端均不存在返回各自错误", clientMissing: true, serverMissing: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				switch r.URL.Path {
				case "/v1.48/containers/json":
					filter, err := filters.FromJSON(r.URL.Query().Get("filters"))
					if err != nil || len(filter.Get("name")) != 1 {
						t.Errorf("unexpected container filter: %s", r.URL.RawQuery)
						w.WriteHeader(http.StatusBadRequest)
						return
					}
					name := filter.Get("name")[0]
					if name != "client_robot-test" && name != "server_robot-test" {
						t.Errorf("unexpected container name: %s", name)
					}
					if (name == "client_robot-test" && tc.clientMissing) || (name == "server_robot-test" && tc.serverMissing) {
						_, _ = fmt.Fprint(w, `[]`)
						return
					}
					_ = json.NewEncoder(w).Encode([]map[string]string{{"Id": name}})
				case "/v1.48/containers/client_robot-test/logs":
					if tc.clientFailure {
						w.WriteHeader(http.StatusInternalServerError)
						_, _ = fmt.Fprint(w, `{"message":"log driver unavailable"}`)
						return
					}
					_, _ = fmt.Fprintln(w, "client startup failed")
				case "/v1.48/containers/server_robot-test/logs":
					_, _ = fmt.Fprintln(w, "server startup complete")
				default:
					t.Errorf("unexpected Docker request: %s %s", r.Method, r.URL)
					w.WriteHeader(http.StatusNotFound)
				}
			}))
			defer server.Close()
			t.Setenv("DOCKER_HOST", strings.Replace(server.URL, "http://", "tcp://", 1))
			t.Setenv("DOCKER_API_VERSION", "1.48")
			t.Setenv("DOCKER_TLS_VERIFY", "")
			t.Setenv("DOCKER_CERT_PATH", "")

			logs, err := NewDockerService(context.Background()).GetRobotContainerLogs(&model.Robot{RobotCode: "robot-test"})
			if err != nil {
				t.Fatalf("GetRobotContainerLogs() error = %v", err)
			}
			if (logs.ClientError != "") != (tc.clientMissing || tc.clientFailure) || (logs.ServerError != "") != tc.serverMissing {
				t.Fatalf("unexpected per-container errors: %+v", logs)
			}
			if !tc.clientMissing && !tc.clientFailure && strings.Join(logs.Client, "\n") != "client startup failed" {
				t.Fatalf("client logs were lost: %+v", logs)
			}
			if !tc.serverMissing && strings.Join(logs.Server, "\n") != "server startup complete" {
				t.Fatalf("server logs were lost: %+v", logs)
			}
		})
	}
}
