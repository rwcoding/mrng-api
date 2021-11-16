package gw

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id     uint32 `validate:"required,numeric,min=1" json:"id"`
	Name   string `validate:"required,max=100" json:"name"`
	Api    string `validate:"required,max=200" json:"api"`
	Key    string `validate:"omitempty,max=200" json:"key"`
	Weight uint8  `validate:"omitempty,numeric,max=100,min=1" json:"weight"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {

	p := models.Gw{}
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的网关")
	}

	isNeedSync := p.Weight == request.Weight

	p.Name = strings.TrimSpace(request.Name)
	p.Api = strings.TrimSpace(request.Api)
	p.Key = strings.TrimSpace(request.Key)
	p.Weight = request.Weight

	if db.Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("修改失败")
	}
	if isNeedSync {
		//todo 同步配置中心
	}
	return api.NewMDResponse("修改成功", &addResponse{
		Id: p.Id,
	})
}
