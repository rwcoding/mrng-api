package white

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id uint32 `validate:"required,numeric,min=1" json:"id"`
	Ip string `validate:"required,max=50" json:"ip"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {

	p := models.ConfigWhite{}
	if db().Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的ip")
	}

	p.Ip = strings.TrimSpace(request.Ip)

	if db().Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("修改失败")
	}

	err := services.SetCacheForWhite(p.Ip)
	if err != nil {
		return api.NewMDResponse("修改成功但缓存失败："+err.Error(), &addResponse{
			Id: p.Id,
		})
	}

	return api.NewMDResponse("修改成功", &addResponse{
		Id: p.Id,
	})
}
