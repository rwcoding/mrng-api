package gw

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
)

type bindQueryRequest struct {
	ctx *boot.Context

	Page     int    `validate:"required,numeric,min=1" json:"page"`
	PageSize int    `validate:"omitempty,numeric,max=10000" json:"page_size"`
	GwId     int    `validate:"required,numeric" json:"gw_id"`
	Name     string `validate:"omitempty,max=100" json:"name"`
	Addr     string `validate:"omitempty,max=100" json:"addr"`
	Bind     int    `validate:"omitempty,numeric,min=0,max=2" json:"bind"`
}

type nodeItemResponse struct {
	Id     uint32 `json:"id"`
	Name   string `json:"name"`
	Addr   string `json:"addr"`
	Status uint8  `json:"status"`
	Weight uint8  `json:"weight"`
}

type bindQueryListResponse struct {
	Datas    []nodeItemResponse `json:"datas"`
	Binds    []uint32           `json:"binds"`
	Count    int64              `json:"count"`
	Page     int                `json:"page"`
	PageSize int                `json:"page_size"`
}

func NewApiBindQuery(ctx *boot.Context) boot.Logic {
	return &bindQueryRequest{ctx: ctx}
}

func (request *bindQueryRequest) Run() *api.Response {
	p := models.Gw{}
	if db().Take(&p, request.GwId).Error != nil {
		return api.NewErrorResponse("无效的网关")
	}

	var bs []models.GwNode
	bindNodeIds := []uint32{}
	db().Where("gw_id = ?", p.Id).Find(&bs)
	for _, v := range bs {
		bindNodeIds = append(bindNodeIds, v.NodeId)
	}

	pageSize := request.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (request.Page - 1) * pageSize
	var us []models.Node
	var c int64

	tx1 := db().Order("id desc").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.Node{})
	if request.Bind == 1 {
		tx1.Where("id IN ?", bindNodeIds)
		tx2.Where("id IN ?", bindNodeIds)
	}
	if request.Bind == 2 {
		tx1.Where("id NOT IN ?", bindNodeIds)
		tx2.Where("id NOT IN ?", bindNodeIds)
	}
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

	list := []nodeItemResponse{}
	for _, v := range us {
		list = append(list, nodeItemResponse{
			Id:     v.Id,
			Name:   v.Name,
			Addr:   v.Addr,
			Status: v.Status,
			Weight: v.Weight,
		})
	}

	return api.NewDataResponse(&bindQueryListResponse{
		Datas:    list,
		Binds:    bindNodeIds,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
