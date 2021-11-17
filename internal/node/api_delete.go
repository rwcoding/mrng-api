package node

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
	var m models.Node
	if db().Take(&m, request.Id).Error != nil {
		return api.NewErrorResponse("无效的节点")
	}

	if db().Delete(&m).RowsAffected == 0 {
		return api.NewErrorResponse("删除失败")
	}

	db().Where("node_id=?", m.Id).Delete(&models.GwNode{})
	db().Where("node_id=?", m.Id).Delete(&models.NodeService{})

	go (func() { services.SyncNodeDelete(m) })()

	return api.NewSuccessResponse("删除成功")
}
