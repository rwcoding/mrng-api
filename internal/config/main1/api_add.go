package main1

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"strings"
)

type addRequest struct {
	ctx *boot.Context

	Name string `validate:"required,max=100" json:"name"`
	//Sign    string `validate:"required,max=100" json:"sign"`
	Env     string `validate:"required,max=100" json:"env"`
	Project string `validate:"required,max=100" json:"project"`
	K       string `validate:"required,max=100" json:"k"`
	V       string `validate:"omitempty,max=3000" json:"v"`
	Status  uint8  `validate:"omitempty,numeric,min=0,max=1" json:"status"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
	p := models.ConfigMain{
		Name: strings.TrimSpace(request.Name),
		//Sign:    strings.TrimSpace(request.Sign),
		Env:     strings.TrimSpace(request.Env),
		Project: strings.TrimSpace(request.Project),
		K:       strings.TrimSpace(request.K),
		V:       strings.TrimSpace(request.V),
		//Status:  request.Status,
		Status: 0,
	}

	p.Sign = p.Env + "." + p.Project + "." + p.K

	if db.Create(&p).RowsAffected == 0 {
		return api.NewErrorResponse("添加失败")
	}

	// 日志
	db.Create(&models.ConfigLog{
		Type:      models.LOG_TYPE_CREATE,
		Name:      p.Name,
		Sign:      p.Sign,
		Env:       p.Env,
		K:         p.K,
		V:         p.V,
		Project:   p.Project,
		AdminerId: request.ctx.GetAdminer().Id,
	})

	return api.NewMDResponse("添加成功", &addResponse{
		Id: p.Id,
	})
}
