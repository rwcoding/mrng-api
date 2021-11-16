package env

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"strings"
)

type editRequest struct {
	ctx *boot.Context

	Id    uint32 `validate:"required,numeric,min=1" json:"id"`
	Name  string `validate:"required,max=100" json:"name"`
	KeyV1 string `validate:"required,max=300" json:"key_v1"`
	KeyV2 string `validate:"required,max=300" json:"key_v2"`
	Ord   uint32 `validate:"required,numeric,min=1" json:"ord"`
}

func NewApiEdit(ctx *boot.Context) boot.Logic {
	return &editRequest{ctx: ctx}
}

func (request *editRequest) Run() *api.Response {

	p := models.ConfigEnv{}
	if db.Take(&p, request.Id).Error != nil {
		return api.NewErrorResponse("无效的环境")
	}

	p.Name = strings.TrimSpace(request.Name)
	p.KeyV1 = request.KeyV1
	p.KeyV2 = request.KeyV2
	p.Ord = request.Ord

	if db.Save(&p).RowsAffected == 0 {
		return api.NewErrorResponse("修改失败")
	}

	err := services.SetCacheForEnv(p.Sign, p.KeyV1, p.KeyV2)
	if err != nil {
		return api.NewMDResponse("添加成功但缓存更新失败："+err.Error(), &addResponse{
			Id: p.Id,
		})
	}

	return api.NewMDResponse("修改成功", &addResponse{
		Id: p.Id,
	})
}
