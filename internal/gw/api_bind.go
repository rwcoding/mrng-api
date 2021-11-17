package gw

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
)

type bindRequest struct {
	ctx *boot.Context

	GwId   int   `validate:"required,numeric" json:"gw_id"`
	Nodes  []int `validate:"required,max=100" json:"nodes"`
	IsBind int   `validate:"omitempty,min=0,max=1" json:"is_bind"`
}

type bindResponse struct {
	Insert int `json:"insert"`
	Delete int `json:"delete"`
}

func NewApiBind(ctx *boot.Context) boot.Logic {
	return &bindRequest{ctx: ctx}
}

func (request *bindRequest) Run() *api.Response {
	p := models.Gw{}
	if db().Take(&p, request.GwId).Error != nil {
		return api.NewErrorResponse("无效的网关")
	}

	var gwNodes []models.GwNode
	db().Where("gw_id=?", p.Id).Find(&gwNodes)

	//var nodes []models.Node
	//if len(request.Nodes) > 0 {
	//	db().Find(&nodes, request.Nodes)
	//}

	var insert []models.GwNode
	isBind := request.IsBind == 1
	ic := 0
	dc := 0
	for _, v := range request.Nodes {
		in := false
		for _, vv := range gwNodes {
			if int(vv.NodeId) == v {
				in = true
				if !isBind {
					db().Delete(&models.GwNode{}, vv.Id)
					dc += 1
				}
			}
		}
		if in || !isBind {
			continue
		}
		//insert
		insert = append(insert, models.GwNode{
			GwId:   p.Id,
			NodeId: uint32(v),
		})
	}
	if len(insert) > 0 {
		ret := db().Create(insert)
		ic = int(ret.RowsAffected)
	}

	if ic > 0 || dc > 0 {
		go (func() { services.SyncGw(p) })()
	}

	return api.NewDataResponse(&bindResponse{
		Insert: ic,
		Delete: dc,
	})
}
