package kv

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
	K         string `json:"k"`
	V         string `json:"v"`
	CreatedAt uint32 `json:"created_at"`
	UpdatedAt uint32 `json:"updated_at"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	var p models.ConfigKv
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的数据")
	}
	return api.NewDataResponse(&infoResponse{
		Id:        p.Id,
		K:         p.K,
		V:         p.V,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
}
