package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"wechat-robot-admin-backend/middleware"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/vars"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestExpiredSessionReturnsConfiguredLoginMethod(t *testing.T) {
	gin.SetMode(gin.TestMode)
	originalLoginMethod := vars.LoginMethod
	t.Cleanup(func() { vars.LoginMethod = originalLoginMethod })

	for _, method := range []string{"token", "scan"} {
		t.Run(method, func(t *testing.T) {
			vars.LoginMethod = method
			router := gin.New()
			router.Use(sessions.Sessions("session", cookie.NewStore([]byte("login-method-test-secret"))))
			router.GET("/user/self", NewUserController().LoginUser)
			router.DELETE("/user/logout", NewUserController().Logout)
			for _, auth := range []struct {
				path    string
				handler gin.HandlerFunc
			}{
				{"/user", middleware.UserAuth()},
				{"/admin", middleware.AdminAuth()},
				{"/root", middleware.RootAuth()},
				{"/owner", middleware.UserOwnerAuth()},
			} {
				router.GET(auth.path, auth.handler, func(c *gin.Context) {
					t.Error("expired session reached protected handler")
				})
			}

			for _, path := range []string{"/user/self", "/user", "/admin", "/root", "/owner", "/user/logout"} {
				t.Run(path, func(t *testing.T) {
					httpMethod, code := http.MethodGet, 401
					if path == "/user/logout" {
						httpMethod, code = http.MethodDelete, 200
					}
					recorder := httptest.NewRecorder()
					// 浏览器在 Session Cookie 过期后不再发送它。
					router.ServeHTTP(recorder, httptest.NewRequest(httpMethod, path, nil))
					assertLoginMethodResponse(t, recorder, code, method)
				})
			}
		})
	}
}

func TestLoginUserReturnsConfiguredLoginMethodWithValidSession(t *testing.T) {
	gin.SetMode(gin.TestMode)
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
	if err := db.AutoMigrate(&model.User{}); err != nil {
		t.Fatal(err)
	}
	user := &model.User{WeChatId: "test-user", DisplayName: "test", Role: vars.RoleRootUser, Status: vars.UserStatusEnabled}
	if err := db.Create(user).Error; err != nil {
		t.Fatal(err)
	}
	originalDB, originalLoginMethod := vars.DB, vars.LoginMethod
	vars.DB = db
	t.Cleanup(func() { vars.DB, vars.LoginMethod = originalDB, originalLoginMethod })

	router := gin.New()
	router.Use(sessions.Sessions("session", cookie.NewStore([]byte("login-method-test-secret"))))
	router.GET("/user/self", func(c *gin.Context) {
		sessions.Default(c).Set("id", user.ID)
		NewUserController().LoginUser(c)
	})
	for _, method := range []string{"token", "scan"} {
		t.Run(method, func(t *testing.T) {
			vars.LoginMethod = method
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/user/self", nil))
			assertLoginMethodResponse(t, recorder, 200, method)
		})
	}
}

func assertLoginMethodResponse(t *testing.T, recorder *httptest.ResponseRecorder, code int, method string) {
	t.Helper()
	var response struct {
		Code int `json:"code"`
		Data struct {
			LoginMethod string `json:"login_method"`
		} `json:"data"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if recorder.Code != http.StatusOK || response.Code != code || response.Data.LoginMethod != method {
		t.Fatalf("expected HTTP 200, code %d, login_method %q; got HTTP %d: %s", code, method, recorder.Code, recorder.Body.String())
	}
}
