package gw

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
)

type syncRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"omitempty,numeric,min=0" json:"id"`
}

func NewApiSync(ctx *boot.Context) boot.Logic {
	return &syncRequest{ctx: ctx}
}

func (request *syncRequest) Run() *api.Response {
	if request.Id > 0 {
		var m models.Gw
		if db().Take(&m, request.Id).Error != nil {
			return api.NewErrorResponse("无效的网关")
		}
		go (func() {
			services.SyncGw(m)
		})()
	} else {
		go (func() {
			services.Sync()
		})()
	}
	return api.NewSuccessResponse("已开始同步任务")
}
