package white

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
)

type listRequest struct {
	ctx *boot.Context

	Page     int    `validate:"required,numeric,min=1" json:"page"`
	PageSize int    `validate:"omitempty,numeric,max=20" json:"page_size"`
	Ip       string `validate:"omitempty,max=100" json:"ip"`
}

type itemResponse struct {
	Id uint32 `json:"id"`
	Ip string `json:"ip"`
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
	var us []models.ConfigWhite
	var c int64

	tx1 := db.Order("id desc").Offset(offset).Limit(pageSize)
	tx2 := db.Model(&models.ConfigWhite{})
	if request.Ip != "" {
		tx1.Where("ip LIKE ?", "%"+request.Ip+"%")
		tx2.Where("ip LIKE ?", "%"+request.Ip+"%")
	}
	tx1.Find(&us)
	tx2.Count(&c)

	var list []itemResponse
	for _, v := range us {
		list = append(list, itemResponse{
			Id: v.Id,
			Ip: v.Ip,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:    list,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
