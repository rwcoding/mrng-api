package services

import (
	"github.com/rwcoding/mrng/models"
	"strings"
)

type reqSyncUpdate struct {
	Cmd  string  `json:"cmd"`
	Data reqNode `json:"data"`
}

func SyncNodeUpdate(node models.Node) {
	var nodeServices []models.NodeService
	var serviceIdList []uint32
	var services []models.Service
	db().Where("node_id = ?", node.Id).Find(&nodeServices)
	for _, v := range nodeServices {
		serviceIdList = append(serviceIdList, v.ServiceId)
	}
	if len(serviceIdList) > 0 {
		serviceIdList = Unique(serviceIdList)
		db().Find(&services, serviceIdList)
	}

	serverNameList := []string{}
	for _, v := range nodeServices {
		if v.NodeId == node.Id {
			for _, vv := range services {
				if vv.Id == v.ServiceId {
					serverNameList = append(serverNameList, vv.Sign)
				}
			}
		}
	}
	serverNameList = UniqueString(serverNameList)
	reqData := reqSyncUpdate{
		Cmd: "update",
		Data: reqNode{
			Addr:    node.Addr,
			Service: strings.Join(serverNameList, ","),
			Weight:  int(node.Weight),
		},
	}

	var gwNodes []models.GwNode
	db().Where("node_id=?", node.Id).Find(&gwNodes)
	for _, v := range gwNodes {
		var gw models.Gw
		db().Take(&gw, v.GwId)
		if gw.Id > 0 {
			apiRequest(gw.Api, gw.Key, reqData)
		}
	}
}
