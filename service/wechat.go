package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/repository"
	"wechat-robot-admin-backend/utils"
	"wechat-robot-admin-backend/vars"
)

type WeChatService struct {
	ctx context.Context
}

type WeChatServerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func NewWeChatService(ctx context.Context) *WeChatService {
	return &WeChatService{
		ctx: ctx,
	}
}

func (sv *WeChatService) GetWeChatIdByCode(code string) (string, error) {
	if code == "" {
		return "", errors.New("无效的参数")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/wechat/user?code=%s", vars.WeChatServerAddress, code), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", vars.WeChatServerToken)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	httpResponse, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer httpResponse.Body.Close()
	var resp WeChatServerResponse
	err = json.NewDecoder(httpResponse.Body).Decode(&resp)
	if err != nil {
		return "", err
	}
	if !resp.Success {
		return "", errors.New(resp.Message)
	}
	if resp.Data == "" {
		return "", errors.New("验证码错误或已过期")
	}
	return resp.Data, nil
}

func (sv *WeChatService) WeChatAuth(ctx context.Context, code string) (*model.User, error) {
	wechatId, err := sv.GetWeChatIdByCode(code)
	if err != nil {
		return nil, err
	}
	user := repository.NewUserRepo(ctx, vars.DB).GetUserByWeChatID(wechatId)
	if user == nil {
		user = &model.User{
			WeChatId:    wechatId,
			DisplayName: fmt.Sprintf("微信用户-%s", utils.GetRandomString(4)),
			Role:        vars.RoleCommonUser,
			Status:      vars.UserStatusEnabled,
			AvatarUrl:   vars.UserDefaultAvatar,
			LastLoginAt: time.Now().Unix(),
			CreatedAt:   time.Now().Unix(),
		}
		repository.NewUserRepo(ctx, vars.DB).Create(user)
	} else {
		user.LastLoginAt = time.Now().Unix()
		repository.NewUserRepo(ctx, vars.DB).Update(user)
	}
	return user, nil
}
