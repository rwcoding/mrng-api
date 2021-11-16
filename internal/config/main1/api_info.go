package main1

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
)

type infoRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

type infoResponse struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	Sign      string `json:"sign"`
	Env       string `json:"env"`
	Project   string `json:"project"`
	K         string `json:"k"`
	V         string `json:"v"`
	Status    uint8  `json:"status"`
	CreatedAt uint32 `json:"created_at"`
	UpdatedAt uint32 `json:"updated_at"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	var p models.ConfigMain
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的配置")
	}
	return api.NewDataResponse(&infoResponse{
		Id:        p.Id,
		Name:      p.Name,
		Sign:      p.Sign,
		Env:       p.Env,
		Project:   p.Project,
		K:         p.K,
		V:         p.V,
		Status:    p.Status,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
}
