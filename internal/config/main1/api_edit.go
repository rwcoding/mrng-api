package main1

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id      uint32 `validate:"required,numeric,min=1" json:"id"`
	Name    string `validate:"required,max=100" json:"name"`
	Env     string `validate:"required,max=100" json:"env"`
	Project string `validate:"required,max=100" json:"project"`
	V       string `validate:"omitempty,max=3000" json:"v"`
	//K       string `validate:"required,max=100" json:"k"`
	//Status  uint8  `validate:"omitempty,numeric,min=0,max=1" json:"status"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {
	p := models.ConfigMain{}
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的配置")
	}

	p.Name = strings.TrimSpace(request.Name)
	p.Env = request.Env
	p.Project = request.Project
	p.V = request.V
	p.Status = 0
	//p.K = request.K
	//p.Status = request.Status

	if db.Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("修改失败")
	}

	db.Create(&models.ConfigLog{
		Type:      models.LOG_TYPE_UPDATE,
		Name:      p.Name,
		Sign:      p.Sign,
		Env:       p.Env,
		K:         p.K,
		V:         p.V,
		Project:   p.Project,
		AdminerId: request.ctx.GetAdminer().Id,
	})

	//修改配置发布后才会更新缓存

	return api.NewMDResponse("修改成功", &addResponse{
		Id: p.Id,
	})
}
