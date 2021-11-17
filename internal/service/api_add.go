package service

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"strings"
)

type addRequest struct {
	ctx *boot.Context

	Name   string `validate:"required,max=100" json:"name"`
	Sign   string `validate:"required,max=200" json:"sign"`
	Status uint8  `validate:"omitempty,numeric,max=100" json:"status"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
	var count int64
	db().Model(&models.Service{}).Where("sign=?", request.Sign).Count(&count)
	if count > 0 {
		return api.NewErrorResponse("相同标识 " + request.Sign + " 已经存在")
	}

	p := models.Service{
		Name:   strings.TrimSpace(request.Name),
		Sign:   strings.TrimSpace(request.Sign),
		Status: request.Status,
	}

	tx := db().Create(&p)

	if tx.Error != nil {
		return api.NewErrorResponse("添加失败 " + tx.Error.Error())
	}

	return api.NewMDResponse("添加成功", &addResponse{
		Id: p.Id,
	})
}
