package node

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
	Addr     string `validate:"omitempty,max=100" json:"addr"`
}

type itemResponse struct {
	Id       uint32 `json:"id"`
	Name     string `json:"name"`
	Addr     string `json:"addr"`
	Services string `json:"services"`
	Status   uint8  `json:"status"`
	Weight   uint8  `json:"weight"`
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
	var us []models.Node
	var c int64

	tx1 := db().Order("id desc").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.Node{})
	if request.Name != "" {
		tx1.Where("name LIKE ?", "%"+request.Name+"%")
		tx2.Where("name LIKE ?", "%"+request.Name+"%")
	}
	if request.Addr != "" {
		tx1.Where("addr LIKE ?", "%"+request.Addr+"%")
		tx2.Where("addr LIKE ?", "%"+request.Addr+"%")
	}
	tx1.Find(&us)
	tx2.Count(&c)

	var list []itemResponse
	for _, v := range us {
		list = append(list, itemResponse{
			Id:       v.Id,
			Name:     v.Name,
			Addr:     v.Addr,
			Services: v.Services,
			Status:   v.Status,
			Weight:   v.Weight,
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:    list,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
