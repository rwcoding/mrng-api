package white

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"strings"
)

type addRequest struct {
	ctx *boot.Context

	Ip string `validate:"required,max=100" json:"ip"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
	p := models.ConfigWhite{
		Ip: strings.TrimSpace(request.Ip),
	}

	if db.Create(&p).RowsAffected == 0 {
		return api.NewErrorResponse("添加失败")
	}

	err := services.SetCacheForWhite(p.Ip)
	if err != nil {
		return api.NewMDResponse("添加成功但缓存失败："+err.Error(), &addResponse{
			Id: p.Id,
		})
	}

	return api.NewMDResponse("添加成功", &addResponse{
		Id: p.Id,
	})
}
