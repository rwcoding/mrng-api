package white

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
)

type infoRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
}

type infoResponse struct {
	Id uint32 `json:"id"`
	Ip string `json:"ip"`
}

func NewApiInfo(ctx *boot.Context) boot.Logic {
	return &infoRequest{ctx: ctx}
}

func (request *infoRequest) Run() *api.Response {
	var p models.ConfigWhite
	if db().Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的数据")
	}
	return api.NewDataResponse(&infoResponse{
		Id: p.Id,
		Ip: p.Ip,
	})
}
