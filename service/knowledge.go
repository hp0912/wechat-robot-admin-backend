package service

import (
	"context"
	"strconv"

	"github.com/go-resty/resty/v2"

	"wechat-robot-admin-backend/dto"
	"wechat-robot-admin-backend/model"
	"wechat-robot-admin-backend/pkg/appx"
)

type KnowledgeService struct {
	ctx context.Context
}

func NewKnowledgeService(ctx context.Context) *KnowledgeService {
	return &KnowledgeService{ctx: ctx}
}

func (s *KnowledgeService) AddDocument(robot *model.Robot, req *dto.AddKnowledgeDocumentRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/knowledge/document")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) UpdateDocument(robot *model.Robot, req *dto.UpdateKnowledgeDocumentRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Put(robot.GetBaseURL() + "/knowledge/document")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) DeleteDocument(robot *model.Robot, req *dto.DeleteKnowledgeRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/knowledge/document")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) EnableDocument(robot *model.Robot, id int64) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]int64{"id": id}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/knowledge/document/enable")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) DisableDocument(robot *model.Robot, id int64) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(map[string]int64{"id": id}).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/knowledge/document/disable")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) ListDocuments(robot *model.Robot, pager appx.Pager, req *dto.ListKnowledgeRequest) (dto.ListResponse[*dto.KnowledgeDocument], error) {
	var result dto.Response[dto.ListResponse[*dto.KnowledgeDocument]]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("category", req.Category).
		SetQueryParam("page_index", strconv.Itoa(pager.PageIndex)).
		SetQueryParam("page_size", "20").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/knowledge/documents")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) GetCategories(robot *model.Robot, req *dto.ListKnowledgeCategoryRequest) ([]*dto.KnowledgeCategory, error) {
	var result dto.Response[[]*dto.KnowledgeCategory]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("type", req.Type).
		SetResult(&result).
		Get(robot.GetBaseURL() + "/knowledge/categories")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) CreateKnowledgeCategory(robot *model.Robot, req *dto.CreateKnowledgeCategoryRequest) (*dto.KnowledgeCategory, error) {
	var result dto.Response[*dto.KnowledgeCategory]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/knowledge/category")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) UpdateKnowledgeCategory(robot *model.Robot, req *dto.UpdateKnowledgeCategoryRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Put(robot.GetBaseURL() + "/knowledge/category")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) DeleteKnowledgeCategory(robot *model.Robot, req *dto.DeleteKnowledgeCategoryRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/knowledge/category")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) SearchKnowledge(robot *model.Robot, req *dto.SearchKnowledgeRequest) ([]dto.VectorSearchResult, error) {
	var result dto.Response[[]dto.VectorSearchResult]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/knowledge/search")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) ReindexAll(robot *model.Robot) (string, error) {
	var result dto.Response[string]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Post(robot.GetBaseURL() + "/knowledge/reindex")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) SaveMemory(robot *model.Robot, req *dto.SaveMemoryRequest) (*dto.Memory, error) {
	var result dto.Response[*dto.Memory]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/memory")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) SearchMemory(robot *model.Robot, req *dto.SearchMemoryRequest) ([]*dto.Memory, error) {
	var result dto.Response[[]*dto.Memory]
	formData := map[string]string{}
	if req.ContactWxID != "" {
		formData["contact_wxid"] = req.ContactWxID
	}
	if req.Query != "" {
		formData["query"] = req.Query
	}
	if req.Limit > 0 {
		formData["limit"] = strconv.Itoa(req.Limit)
	}
	request := resty.New().R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetResult(&result)
	if len(formData) > 0 {
		request.SetFormData(formData)
	}
	_, err := request.Post(robot.GetBaseURL() + "/memory/search")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) DeleteMemory(robot *model.Robot, req *dto.DeleteMemoryRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/memory")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) AddImageDocument(robot *model.Robot, req *dto.AddImageKnowledgeRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/image-knowledge/document")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) DeleteImageDocument(robot *model.Robot, req *dto.DeleteImageKnowledgeRequest) error {
	var result dto.Response[struct{}]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Delete(robot.GetBaseURL() + "/image-knowledge/document")
	if err = result.CheckError(err); err != nil {
		return err
	}
	return nil
}

func (s *KnowledgeService) ListImageDocuments(robot *model.Robot, pager appx.Pager, req *dto.ListImageKnowledgeRequest) (dto.ListResponse[*dto.ImageKnowledgeDocument], error) {
	var result dto.Response[dto.ListResponse[*dto.ImageKnowledgeDocument]]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetQueryParam("category", req.Category).
		SetQueryParam("page_index", strconv.Itoa(pager.PageIndex)).
		SetQueryParam("page_size", "20").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/image-knowledge/documents")

	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) GetImageCategories(robot *model.Robot) ([]string, error) {
	var result dto.Response[[]string]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Get(robot.GetBaseURL() + "/image-knowledge/categories")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) SearchImageByText(robot *model.Robot, req *dto.SearchImageKnowledgeByTextRequest) ([]dto.VectorSearchResult, error) {
	var result dto.Response[[]dto.VectorSearchResult]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/image-knowledge/search/text")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) SearchImageByImage(robot *model.Robot, req *dto.SearchImageKnowledgeByImageRequest) ([]dto.VectorSearchResult, error) {
	var result dto.Response[[]dto.VectorSearchResult]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetBody(req).
		SetResult(&result).
		Post(robot.GetBaseURL() + "/image-knowledge/search/image")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}

func (s *KnowledgeService) ReindexAllImages(robot *model.Robot) (string, error) {
	var result dto.Response[string]
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json;chartset=utf-8").
		SetResult(&result).
		Post(robot.GetBaseURL() + "/image-knowledge/reindex")
	if err = result.CheckError(err); err != nil {
		return result.Data, err
	}
	return result.Data, nil
}
