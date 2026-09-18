package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRobotCreateFailureReturnsPersistedInstanceID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	for _, tc := range []struct {
		name    string
		version string
		wantID  bool
	}{
		{name: "初始化失败保留实例ID", version: "8.0.59", wantID: true},
		{name: "实例创建前失败没有ID", version: "unsupported"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
			if err != nil {
				t.Fatal(err)
			}
			sqlDB, err := db.DB()
			if err != nil {
				t.Fatal(err)
			}
			sqlDB.SetMaxOpenConns(1)
			t.Cleanup(func() { _ = sqlDB.Close() })
			// SQLite 用 text 代替 MySQL enum；后续 CREATE DATABASE 会失败，模拟实例落库后的初始化失败。
			if err := db.Table("robot").AutoMigrate(&struct {
				model.Robot
				Status string `gorm:"column:status;type:text"`
			}{}); err != nil {
				t.Fatal(err)
			}
			originalDB := vars.DB
			vars.DB = db
			t.Cleanup(func() { vars.DB = originalDB })

			recorder := httptest.NewRecorder()
			router := gin.New()
			router.Use(sessions.Sessions("test", cookie.NewStore([]byte("robot-create-test-secret"))))
			router.POST("/api/v1/robot/create", func(ctx *gin.Context) {
				session := sessions.Default(ctx)
				session.Set("wechat_id", "test-owner")
				session.Set("role", vars.RoleCommonUser)
				NewRobotManageController().RobotCreate(ctx)
			})
			request := httptest.NewRequest(http.MethodPost, "/api/v1/robot/create", strings.NewReader(`{"robot_name":"test robot","version":"`+tc.version+`"}`))
			request.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(recorder, request)

			var response dto.Response[*dto.RobotCreateResponse]
			if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			if recorder.Code != http.StatusOK || response.Code != 500 || response.Message == "" {
				t.Fatalf("unexpected failure response: %s", recorder.Body.String())
			}
			if !tc.wantID {
				if response.Data != nil {
					t.Fatalf("failure before creation returned an instance: %+v", response.Data)
				}
				return
			}
			if response.Data == nil || response.Data.ID <= 0 {
				t.Fatalf("failure lost the persisted instance ID: %s", recorder.Body.String())
			}
			var robot model.Robot
			if err := db.First(&robot, response.Data.ID).Error; err != nil {
				t.Fatal(err)
			}
			if robot.Owner != "test-owner" || robot.RobotName != "test robot" {
				t.Fatalf("response ID points at another instance: %+v", robot)
			}
			var payload dto.Response[map[string]json.RawMessage]
			if err := json.Unmarshal(recorder.Body.Bytes(), &payload); err != nil {
				t.Fatal(err)
			}
			if len(payload.Data) != 1 {
				t.Fatalf("creation response should expose only the ID: %s", recorder.Body.String())
			}
		})
	}
}
