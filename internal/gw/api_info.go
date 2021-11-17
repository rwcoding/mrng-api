package gw

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
	Addr   string `json:"addr"`
	Api    string `json:"api"`
	Status uint8  `json:"status"`
	Weight uint8  `json:"weight"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	var p models.Gw
	if db().Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的网关")
	}
	return api.NewDataResponse(&infoResponse{
		Id:     p.Id,
		Name:   p.Name,
		Addr:   p.Addr,
		Api:    p.Api,
		Status: p.Status,
		Weight: p.Weight,
	})
}
