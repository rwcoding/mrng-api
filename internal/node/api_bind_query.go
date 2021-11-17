package node

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
)

type bindQueryRequest struct {
	ctx *boot.Context

	Page     int    `validate:"required,numeric,min=1" json:"page"`
	PageSize int    `validate:"omitempty,numeric,max=10000" json:"page_size"`
	NodeId   int    `validate:"required,numeric" json:"node_id"`
	Name     string `validate:"omitempty,max=100" json:"name"`
	Sign     string `validate:"omitempty,max=100" json:"sign"`
	Bind     int    `validate:"omitempty,numeric,min=0,max=2" json:"bind"`
}

type serviceItemResponse struct {
	Id     uint32 `json:"id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Status uint8  `json:"status"`
}

type bindQueryListResponse struct {
	Datas    []serviceItemResponse `json:"datas"`
	Binds    []uint32              `json:"binds"`
	Count    int64                 `json:"count"`
	Page     int                   `json:"page"`
	PageSize int                   `json:"page_size"`
}

func NewApiBindQuery(ctx *boot.Context) boot.Logic {
	return &bindQueryRequest{ctx: ctx}
}

func (request *bindQueryRequest) Run() *api.Response {
	p := models.Node{}
	if db().Take(&p, request.NodeId).Error != nil {
		return api.NewErrorResponse("无效的节点")
	}

	var bs []models.NodeService
	bindIds := []uint32{}
	db().Where("node_id = ?", p.Id).Find(&bs)
	for _, v := range bs {
		bindIds = append(bindIds, v.ServiceId)
	}

	pageSize := request.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (request.Page - 1) * pageSize
	var us []models.Service
	var c int64

	tx1 := db().Order("id desc").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.Service{})
	if request.Bind == 1 {
		tx1.Where("id IN ?", bindIds)
		tx2.Where("id IN ?", bindIds)
	}
	if request.Bind == 2 {
		tx1.Where("id NOT IN ?", bindIds)
		tx2.Where("id NOT IN ?", bindIds)
	}
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

	list := []serviceItemResponse{}
	for _, v := range us {
		list = append(list, serviceItemResponse{
			Id:     v.Id,
			Name:   v.Name,
			Sign:   v.Sign,
			Status: v.Status,
		})
	}

	return api.NewDataResponse(&bindQueryListResponse{
		Datas:    list,
		Binds:    bindIds,
		Count:    c,
		Page:     request.Page,
		PageSize: pageSize,
	})
}
