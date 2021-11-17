package env

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"strings"
)

type addRequest struct {
	ctx *boot.Context

	Name  string `validate:"required,max=100" json:"name"`
	Sign  string `validate:"required,max=100" json:"sign"`
	KeyV1 string `validate:"required,max=300" json:"key_v1"`
	KeyV2 string `validate:"required,max=300" json:"key_v2"`
	Ord   uint32 `validate:"required,numeric,min=1" json:"ord"`
}

type addResponse struct {
	Id uint32 `json:"id"`
}

func NewApiAdd(ctx *boot.Context) boot.Logic {
	return &addRequest{ctx: ctx}
}

func (request *addRequest) Run() *api.Response {
	p := models.ConfigEnv{
		Name:  strings.TrimSpace(request.Name),
		Sign:  strings.TrimSpace(request.Sign),
		KeyV1: strings.TrimSpace(request.KeyV1),
		KeyV2: strings.TrimSpace(request.KeyV2),
		Ord:   request.Ord,
	}

	if strings.Contains(p.Sign, ".") || strings.Contains(p.Sign, " ") {
		return api.NewErrorResponse("标识中不允许空白、点等特殊字符")
	}

	if db().Create(&p).RowsAffected == 0 {
		return api.NewErrorResponse("添加失败")
	}

	err := services.SetCacheForEnv(p.Sign, p.KeyV1, p.KeyV2)
	if err != nil {
		return api.NewMDResponse("添加成功但缓存更新失败："+err.Error(), &addResponse{
			Id: p.Id,
		})
	}

	return api.NewMDResponse("添加成功", &addResponse{
		Id: p.Id,
	})
}
