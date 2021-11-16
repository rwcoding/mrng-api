package node

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id   uint32 `validate:"required,numeric,min=1" json:"id"`
	Name string `validate:"required,max=100" json:"name"`
	//Addr   string `validate:"required,max=200" json:"addr"`
	Status uint8 `validate:"omitempty,numeric,max=100,min=0" json:"status"`
	Weight uint8 `validate:"omitempty,numeric,max=100,min=0" json:"weight"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {

	p := models.Node{}
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的节点")
	}

	p.Name = strings.TrimSpace(request.Name)
	//p.Addr = strings.TrimSpace(request.Addr)
	p.Status = request.Status
	p.Weight = request.Weight

	if db.Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("修改失败")
	}

	go (func() { services.SyncNodeUpdate(p) })()

	return api.NewMDResponse("修改成功", &addResponse{
		Id: p.Id,
	})
}
