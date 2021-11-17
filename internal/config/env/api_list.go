package env

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
)

type listRequest struct {
	ctx *boot.Context

	Page     int    `validate:"required,numeric,min=1" json:"page"`
	PageSize int    `validate:"omitempty,numeric,max=20" json:"page_size"`
	Name     string `validate:"omitempty,max=100" json:"name"`
	Sign     string `validate:"omitempty,max=100" json:"sign"`
}

type itemResponse struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Sign    string `json:"sign"`
	KeyV1   string `json:"key_v1"`
	KeyV2   string `json:"key_v2"`
	Ord     uint32 `json:"ord"`
	Version int64  `json:"version"`
}

type listResponse struct {
	Datas    []itemResponse `json:"datas"`
	Count    int64          `json:"count"`
	Page     int            `json:"page"`
	PageSize int            `json:"page_size"`
}

func NewApiList(ctx *boot.Context) boot.Logic {
	return &listRequest{ctx: ctx}
}

func (request *listRequest) Run() *api.Response {
	pageSize := request.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (request.Page - 1) * pageSize
	var us []models.ConfigEnv
	var c int64

	tx1 := db().Order("ord").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.ConfigEnv{})
	if request.Name != "" {
		tx1.Where("name LIKE ?", "%"+request.Name+"%")
		tx2.Where("name LIKE ?", "%"+request.Name+"%")
	}
	if request.Sign != "" {
		tx1.Where("sign LIKE ?", "%"+request.Sign+"%")
		tx2.Where("sign LIKE ?", "%"+request.Sign+"%")
	}
	tx1.Find(&us)
	tx2.Count(&c)

	var list []itemResponse
	for _, v := range us {
		list = append(list, itemResponse{
			Id:      v.Id,
			Name:    v.Name,
			Sign:    v.Sign,
			KeyV1:   v.KeyV1,
			KeyV2:   v.KeyV2,
			Ord:     v.Ord,
			Version: v.Version,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:    list,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
