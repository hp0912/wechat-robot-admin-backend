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

func (sv *WeChatService) WeChatOfficialAccountAuthURL(ctx context.Context) {

}

func (sv *WeChatService) WeChatAuth(ctx context.Context, code string) (*model.User, error) {
	wechatId, err := sv.GetWeChatIdByCode(code)
	if err != nil {
		return nil, err
	}
	userRespo := repository.NewUserRepo(ctx, vars.DB)
	user, err := userRespo.GetUserByWeChatID(wechatId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		userCount, err := userRespo.UserCount()
		if err != nil {
			return nil, fmt.Errorf("获取用户数量失败，请联系管理员: %w", err)
		}
		user = &model.User{
			WeChatId:    wechatId,
			DisplayName: fmt.Sprintf("微信用户-%s", utils.GetRandomString(4)),
			Status:      vars.UserStatusEnabled,
			AvatarUrl:   vars.UserDefaultAvatar,
			LastLoginAt: time.Now().Unix(),
			CreatedAt:   time.Now().Unix(),
		}
		if userCount == 0 {
			user.Role = vars.RoleRootUser // 第一个用户设置为超级管理员
		} else {
			user.Role = vars.RoleCommonUser
		}
		err = userRespo.Create(user)
		if err != nil {
			return nil, fmt.Errorf("登录失败，请联系管理员: %w", err)
		}
	} else {
		user.LastLoginAt = time.Now().Unix()
		err = userRespo.Update(user)
		if err != nil {
			return nil, fmt.Errorf("登录失败，请联系管理员: %w", err)
		}
	}
	return user, nil
}
