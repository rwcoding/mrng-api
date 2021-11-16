package gw

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
)

type deleteRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

func NewApiDelete(ctx *boot.Context) boot.Logic {
	return &deleteRequest{ctx: ctx}
}

func (request *deleteRequest) Run() *api.Response {
	var m models.Gw
	if db.Take(&m, request.Id).Error != nil {
		return api.NewErrorResponse("无效的网关")
	}

	if db.Delete(&m).RowsAffected == 0 {
		return api.NewErrorResponse("删除失败")
	}

	db.Where("gw_id=?", m.Id).Delete(&models.GwNode{})

	//todo 同步配置中心
	return api.NewSuccessResponse("删除成功")
}
