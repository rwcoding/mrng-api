package main1

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
)

type deleteRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

func NewApiDelete(ctx *boot.Context) boot.Logic {
	return &deleteRequest{ctx: ctx}
}

func (request *deleteRequest) Run() *api.Response {
	adminerId := request.ctx.GetAdminer().Id
	var m models.ConfigMain
	if db().Take(&m, request.Id).Error != nil {
		return api.NewErrorResponse("无效的配置")
	}

	if db().Delete(&m).RowsAffected == 0 {
		return api.NewErrorResponse("删除失败")
	}

	db().Create(&models.ConfigLog{
		Type:      models.LOG_TYPE_DELETE,
		Name:      m.Name,
		Sign:      m.Sign,
		Env:       m.Env,
		K:         m.K,
		V:         m.V,
		Project:   m.Project,
		AdminerId: adminerId,
	})

	err := services.DeleteCacheForConfig(m.Sign)
	if err != nil {
		return api.NewSuccessResponse("删除成功但未删除缓存：" + err.Error())
	}

	return api.NewSuccessResponse("删除成功")
}
