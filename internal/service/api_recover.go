package service

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
)

type recoverRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

func NewApiRecover(ctx *boot.Context) boot.Logic {
	return &recoverRequest{ctx: ctx}
}

func (request *recoverRequest) Run() *api.Response {
	var u models.Service
	if db.Unscoped().Take(&u, request.Id).Error != nil {
		return api.NewErrorResponse("无效的服务")
	}

	if db.Model(&u).Update("deleted_at", 0).RowsAffected == 0 {
		return api.NewErrorResponse("恢复失败")
	}

	go (func() { services.Sync() })()

	return api.NewSuccessResponse("恢复成功")
}
