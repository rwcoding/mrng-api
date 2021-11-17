package services

import (
	"github.com/rwcoding/mrng/models"
	"strings"
)

type reqCrash struct {
	Crash string `json:"crash"`
}

type reqSyncCrash struct {
	Cmd  string   `json:"cmd"`
	Data reqCrash `json:"data"`
}

func SyncServiceCrash(service models.Service) {
	var nodeServices []models.NodeService
	var gwNodes []models.GwNode
	var nodeIdList []uint32
	var gwIdList []uint32
	var services []models.Service
	var crashList []string
	db().Where("service_id = ?", service.Id).Find(&nodeServices)
	for _, v := range nodeServices {
		nodeIdList = append(nodeIdList, v.NodeId)
	}
	nodeIdList = Unique(nodeIdList)

	if len(nodeIdList) > 0 {
		db().Where("node_id IN ?", nodeIdList).Find(&gwNodes)
		for _, v := range gwNodes {
			gwIdList = append(gwIdList, v.GwId)
		}
		gwIdList = Unique(gwIdList)
	}

	db().Where("status=0").Find(&services)
	for _, v := range services {
		crashList = append(crashList, v.Sign)
	}
	reqData := reqSyncCrash{
		Cmd: "delete",
		Data: reqCrash{
			Crash: strings.Join(crashList, ","),
		},
	}

	if len(gwIdList) > 0 {
		var gws []models.Gw
		db().Find(&gws, gwIdList)
		for _, v := range gws {
			apiRequest(v.Api, v.Key, reqData)
		}
	}
}
