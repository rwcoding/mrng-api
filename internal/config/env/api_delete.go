package env

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
	var m models.ConfigEnv
	if db.Take(&m, request.Id).Error != nil {
		return api.NewErrorResponse("无效的环境")
	}

	if db.Delete(&m).RowsAffected == 0 {
		return api.NewErrorResponse("删除失败")
	}

	db.Where("env = ?", m.Sign).Delete(&models.ConfigMain{})

	err := services.DeleteCacheForEnv(m.Sign)
	if err != nil {
		return api.NewSuccessResponse("删除成功但缓存更新失败：" + err.Error())
	}

	return api.NewSuccessResponse("删除成功")
}
