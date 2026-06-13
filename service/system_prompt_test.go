package service

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
)

func TestListSystemPromptsForwardsKeyword(t *testing.T) {
	var gotPath string
	var gotKeyword string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotKeyword = r.URL.Query().Get("keyword")
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{"code":200,"message":"","data":[{"id":1,"title":"客服","content":"你是客服助手","created_at":1,"updated_at":2}]}`)
	}))
	defer server.Close()

	t.Setenv("DEV_ROBOT_CLIENT_URL", server.URL)
	prompts, err := NewSystemPromptService(context.Background()).ListSystemPrompts(&model.Robot{}, &dto.ListSystemPromptRequest{Keyword: "客服"})
	if err != nil {
		t.Fatalf("ListSystemPrompts returned error: %v", err)
	}
	if gotPath != "/system-prompts" {
		t.Fatalf("path = %q, want %q", gotPath, "/system-prompts")
	}
	if gotKeyword != "客服" {
		t.Fatalf("keyword = %q, want %q", gotKeyword, "客服")
	}
	if len(prompts) != 1 || prompts[0].Title != "客服" || prompts[0].Content != "你是客服助手" {
		t.Fatalf("prompts = %#v", prompts)
	}
}
