package node

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"strings"
)

type bindRequest struct {
	ctx *boot.Context

	NodeId   int   `validate:"required,numeric" json:"node_id"`
	Services []int `validate:"required,max=100" json:"services"`
	IsBind   int   `validate:"omitempty,min=0,max=1" json:"is_bind"`
}

type bindResponse struct {
	Insert int `json:"insert"`
	Delete int `json:"delete"`
}

func NewApiBind(ctx *boot.Context) boot.Logic {
	return &bindRequest{ctx: ctx}
}

func (request *bindRequest) Run() *api.Response {
	p := models.Node{}
	if db.Take(&p, request.NodeId).Error != nil {
		return api.NewErrorResponse("无效的节点")
	}

	var nodeServices []models.NodeService
	db.Where("node_id=?", p.Id).Find(&nodeServices)

	var insert []models.NodeService
	isBind := request.IsBind == 1
	ic := 0
	dc := 0
	for _, v := range request.Services {
		in := false
		for _, vv := range nodeServices {
			if int(vv.ServiceId) == v {
				in = true
				if !isBind {
					db.Delete(&models.NodeService{}, vv.Id)
					dc += 1
				}
			}
		}
		if in || !isBind {
			continue
		}
		//insert
		insert = append(insert, models.NodeService{
			NodeId:    p.Id,
			ServiceId: uint32(v),
		})
	}
	if len(insert) > 0 {
		ret := db.Create(insert)
		ic = int(ret.RowsAffected)
	}

	if ic > 0 || dc > 0 {
		var ns []models.NodeService
		var ids []uint32
		var nameList []string
		db.Where("node_id = ?", p.Id).Find(&ns)
		for _, v := range ns {
			ids = append(ids, v.ServiceId)
		}
		if len(ids) > 0 {
			var services []models.Service
			db.Find(&services, ids)
			for _, v := range services {
				nameList = append(nameList, v.Sign)
			}
		}

		names := ""
		if len(nameList) > 0 {
			names = strings.Join(nameList, ",")
		}
		db.Model(&p).Update("services", names)

		go (func() { services.SyncNodeUpdate(p) })()
	}

	return api.NewDataResponse(&bindResponse{
		Insert: ic,
		Delete: dc,
	})
}
