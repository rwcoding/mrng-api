package gw

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
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

	p := models.Gw{}
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的节点")
	}

	isNeedSync := p.Status != request.Status

	p.Status = request.Status

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
