package services

import (
	"github.com/rwcoding/mrng/models"
	"strings"
)

type reqLower struct {
	Lower string `json:"lower"`
}

type reqSyncLower struct {
	Cmd  string   `json:"cmd"`
	Data reqLower `json:"data"`
}

func SyncServiceLower(service models.Service) {
	var nodeServices []models.NodeService
	var gwNodes []models.GwNode
	var nodeIdList []uint32
	var gwIdList []uint32
	var services []models.Service
	var lowerList []string
	db.Where("service_id = ?", service.Id).Find(&nodeServices)
	for _, v := range nodeServices {
		nodeIdList = append(nodeIdList, v.NodeId)
	}
	nodeIdList = Unique(nodeIdList)

	if len(nodeIdList) > 0 {
		db.Where("node_id IN ?", nodeIdList).Find(&gwNodes)
		for _, v := range gwNodes {
			gwIdList = append(gwIdList, v.GwId)
		}
		gwIdList = Unique(gwIdList)
	}

	db.Where("status>0 AND status<100").Find(&services)
	for _, v := range services {
		lowerList = append(lowerList, v.Sign)
	}
	reqData := reqSyncLower{
		Cmd: "delete",
		Data: reqLower{
			Lower: strings.Join(lowerList, ","),
		},
	}

	if len(gwIdList) > 0 {
		var gws []models.Gw
		db.Find(&gws, gwIdList)
		for _, v := range gws {
			apiRequest(v.Api, v.Key, reqData)
		}
	}
}
