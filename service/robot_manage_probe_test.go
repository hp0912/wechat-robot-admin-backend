package service

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"wechat-robot-admin-backend/model"
)

type robotProbeTransport func(*http.Request) (*http.Response, error)

func (fn robotProbeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func TestWaitRobotClientReady(t *testing.T) {
	tests := []struct {
		name         string
		mode         string
		cancelAfter  time.Duration
		wantDuration time.Duration
		wantErr      error
		wantAttempts int
	}{
		{name: "立即就绪", mode: "ready", wantAttempts: 1},
		{name: "失败响应和网络错误后重试", mode: "retry", wantDuration: 5 * time.Second, wantAttempts: 6},
		{name: "持续未就绪时30秒超时", mode: "pending", wantDuration: 30 * time.Second, wantErr: context.DeadlineExceeded},
		{name: "请求卡住时仍然30秒超时", mode: "stalled", wantDuration: 30 * time.Second, wantErr: context.DeadlineExceeded},
		{name: "轮询等待可以取消", mode: "pending", cancelAfter: 500 * time.Millisecond, wantDuration: 500 * time.Millisecond, wantErr: context.Canceled, wantAttempts: 1},
		{name: "正在请求时可以取消", mode: "stalled", cancelAfter: 500 * time.Millisecond, wantDuration: 500 * time.Millisecond, wantErr: context.Canceled, wantAttempts: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 使用虚拟时间验证完整的 30 秒超时，无需实际等待。
			synctest.Test(t, func(t *testing.T) {
				t.Setenv("DEV_ROBOT_CLIENT_URL", "http://client_robot-test:9000/api/v1/robot")
				originalTransport := http.DefaultTransport
				t.Cleanup(func() {
					http.DefaultTransport = originalTransport
				})

				attempts := 0
				http.DefaultTransport = robotProbeTransport(func(req *http.Request) (*http.Response, error) {
					attempts++
					if req.Method != http.MethodPost || req.URL.Path != "/api/v1/robot/probe" {
						t.Fatalf("unexpected probe request: %s %s", req.Method, req.URL.Path)
					}
					if tt.mode == "stalled" {
						<-req.Context().Done()
						return nil, req.Context().Err()
					}

					status, body := http.StatusOK, `{"success":true}`
					if tt.mode == "pending" {
						body = `{"success":false}`
					}
					if tt.mode == "retry" {
						switch attempts {
						case 1:
							status = http.StatusServiceUnavailable
						case 2:
							body = `{"success":false}`
						case 3:
							body = `invalid json`
						case 4:
							return nil, errors.New("connection refused")
						case 5:
							body = `{}`
						}
					}
					return &http.Response{
						StatusCode: status,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(body)),
						Request:    req,
					}, nil
				})

				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				if tt.cancelAfter > 0 {
					time.AfterFunc(tt.cancelAfter, cancel)
				}
				start := time.Now()
				err := NewRobotManageService(ctx).waitRobotClientReady(&model.Robot{RobotCode: "robot-test"})
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("waitRobotClientReady() error = %v, want %v", err, tt.wantErr)
				}
				if elapsed := time.Since(start); elapsed != tt.wantDuration {
					t.Fatalf("elapsed = %v, want %v", elapsed, tt.wantDuration)
				}
				if tt.wantAttempts > 0 && attempts != tt.wantAttempts {
					t.Fatalf("probe attempts = %d, want %d", attempts, tt.wantAttempts)
				}
			})
		})
	}
}
