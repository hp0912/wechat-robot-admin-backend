package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"wechat-robot-client/model"
	"wechat-robot-client/repository"
	"wechat-robot-client/utils"
	"wechat-robot-client/vars"
)

type WechatService struct {
	ctx context.Context
}

type WechatServerResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func NewWechatService(ctx context.Context) *WechatService {
	return &WechatService{
		ctx: ctx,
	}
}

func (w *WechatService) GetWeChatIdByCode(code string) (string, error) {
	if code == "" {
		return "", errors.New("无效的参数")
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/wechat/user?code=%s", vars.WechatServerAddress, code), nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", vars.WechatServerToken)
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	httpResponse, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer httpResponse.Body.Close()
	var resp WechatServerResponse
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

func (w *WechatService) WechatAuth(ctx context.Context, code string) (*model.User, error) {
	wechatId, err := w.GetWeChatIdByCode(code)
	if err != nil {
		return nil, err
	}
	user := repository.NewUserRepo(ctx, vars.DB).GetUserByWechatID(wechatId)
	if user == nil {
		user = &model.User{
			WeChatId:      wechatId,
			DisplayName:   fmt.Sprintf("微信用户-%s", utils.GetRandomString(4)),
			Role:          vars.RoleCommonUser,
			Status:        vars.UserStatusEnabled,
			AvatarUrl:     vars.UserDefaultAvatar,
			LastLoginTime: time.Now().Unix(),
			CreatedTime:   time.Now().Unix(),
		}
		repository.NewUserRepo(ctx, vars.DB).Create(user)
	} else {
		user.LastLoginTime = time.Now().Unix()
		repository.NewUserRepo(ctx, vars.DB).Update(user)
	}
	return user, nil
}
