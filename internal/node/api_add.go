package node

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"strings"
)

type addRequest struct {
	ctx *boot.Context

	Name   string `validate:"required,max=100" json:"name"`
	Addr   string `validate:"required,max=200" json:"addr"`
	Weight uint8  `validate:"omitempty,numeric,max=100" json:"weight"`
	Status uint8  `validate:"omitempty,numeric,max=100" json:"status"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
	p := models.Node{
		Name:   strings.TrimSpace(request.Name),
		Addr:   strings.TrimSpace(request.Addr),
		Status: request.Status,
		Weight: request.Weight,
	}

	if db().Create(&p).RowsAffected == 0 {
		return api.NewErrorResponse("添加失败")
	}

	return api.NewMDResponse("添加成功", &addResponse{
		Id: p.Id,
	})
}
