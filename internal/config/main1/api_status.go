package main1

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
)

type statusRequest struct {
	ctx *boot.Context

	Id     uint32 `validate:"required,numeric,min=1" json:"id"`
	Status uint8  `validate:"omitempty,numeric,max=1,min=0" json:"status"`
}

func NewApiStatus(ctx *boot.Context) boot.Logic {
	return &statusRequest{ctx: ctx}
}

func (request *statusRequest) Run() *api.Response {

	p := models.ConfigMain{}
	if db().Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的配置")
	}

	p.Status = request.Status

	if db().Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("操作失败")
	}

	if p.Status == 1 {
		err := services.SetCacheForConfig(p.Sign, p.V)
		if err != nil {
			return api.NewMDResponse("发布失败："+err.Error(), &addResponse{
				Id: p.Id,
			})
		}
	} else {
		err := services.DeleteCacheForConfig(p.Sign)
		if err != nil {
			return api.NewMDResponse("状态已修改但未删除缓存："+err.Error(), &addResponse{
				Id: p.Id,
			})
		}
	}

	return api.NewMDResponse("操作成功", &addResponse{
		Id: p.Id,
	})
}
