package service

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
	Min      int    `validate:"omitempty,numeric,min=0,max=100" json:"min"`
	Max      int    `validate:"omitempty,numeric,min=0,max=100" json:"max"`
	Recycle  int    `validate:"omitempty,numeric,min=0,max=1" json:"recycle"`
}

type itemResponse struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Sign    string `json:"sign"`
	Status  uint8  `json:"status"`
	Deleted uint32 `json:"deleted"`
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
	var us []models.Service
	var c int64

	tx1 := db().Order("id desc").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.Service{})
	if request.Recycle == 1 {
		tx1.Unscoped()
		tx2.Unscoped()
	}
	if request.Sign != "" {
		tx1.Where("sign LIKE ?", "%"+request.Sign+"%")
		tx2.Where("sign LIKE ?", "%"+request.Sign+"%")
	}
	if request.Name != "" {
		tx1.Where("name LIKE ?", "%"+request.Name+"%")
		tx2.Where("name LIKE ?", "%"+request.Name+"%")
	}
	if request.Min >= 0 {
		tx1.Where("status >= ?", request.Min)
		tx2.Where("status >= ?", request.Min)
	}
	if request.Max >= 0 {
		tx1.Where("status <= ?", request.Max)
		tx2.Where("status <= ?", request.Max)
	}
	tx1.Find(&us)
	tx2.Count(&c)

	var list []itemResponse
	for _, v := range us {
		list = append(list, itemResponse{
			Id:      v.Id,
			Name:    v.Name,
			Sign:    v.Sign,
			Status:  v.Status,
			Deleted: uint32(v.DeletedAt),
		})
	}

	return api.NewDataResponse(&listResponse{
		Datas:    list,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
