package service

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
	Id     uint32 `json:"id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	var p models.Service
	if db().Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的服务")
	}
	return api.NewDataResponse(&infoResponse{
		Id:     p.Id,
		Name:   p.Name,
		Sign:   p.Sign,
		Status: p.Status,
	})
}
