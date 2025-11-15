package service

import (
	"context"
	"strconv"

	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"

	"github.com/go-resty/resty/v2"
)

type MCPServerService struct {
	ctx context.Context
}

func NewMCPServerService(ctx context.Context) *MCPServerService {
	return &MCPServerService{
		ctx: ctx,
	}
}

func (s *MCPServerService) GetMCPServers(robot *model.Robot) ([]*dto.MCPServer, error) {
	var result dto.Response[[]*dto.MCPServer]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/mcp/servers")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *MCPServerService) GetMCPServer(robot *model.Robot, id int64) (*dto.MCPServer, error) {
	var result dto.Response[*dto.MCPServer]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("id", strconv.FormatInt(id, 10)).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/mcp/server")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *MCPServerService) GetMCPServerTools(robot *model.Robot, id int64) ([]*dto.MCPServerTool, error) {
	var result dto.Response[[]*dto.MCPServerTool]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("id", strconv.FormatInt(id, 10)).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/mcp/server/tools")
	if err = result.CheckError(err); err != nil {
		return nil, err
	}
	return result.Data, nil
}

func (s *MCPServerService) CreateMCPServer(robot *model.Robot, mcpServer *dto.MCPServer) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(mcpServer).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/mcp/server")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *MCPServerService) UpdateMCPServer(robot *model.Robot, mcpServer *dto.MCPServer) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(mcpServer).
		SetResult(&result).
		Put(robot.GetBaseURL() + "/mcp/server")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *MCPServerService) EnableMCPServer(robot *model.Robot, id int64) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]int64{
			"id": id,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/mcp/server/enable")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *MCPServerService) DisableMCPServer(robot *model.Robot, id int64) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]int64{
			"id": id,
		}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/mcp/server/disable")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *MCPServerService) DeleteMCPServer(robot *model.Robot, mcpServer *dto.MCPServer) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(mcpServer).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/mcp/server")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}
