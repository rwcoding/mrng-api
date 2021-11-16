package services

import (
	"github.com/rwcoding/mrng/models"
	"strings"
)

type reqNode struct {
	Addr    string `json:"addr"`
	Service string `json:"service"`
	Weight  int    `json:"weight"`
}

type reqSyncData struct {
	Nodes []reqNode      `json:"nodes"`
	Crash string         `json:"crash"`
	Lower map[string]int `json:"lower"`
}

type reqSync struct {
	Cmd  string      `json:"cmd"`
	Data reqSyncData `json:"data"`
}

func Sync() {
	var gws []models.Gw
	db.Where("status=1").Find(&gws)
	for _, v := range gws {
		SyncGw(v)
	}
}

func SyncGw(gw models.Gw) {
	var gwNodes []models.GwNode
	var nodeIdList []uint32
	var nodes []models.Node
	db.Where("gw_id=?", gw.Id).Find(&gwNodes)
	for _, v := range gwNodes {
		nodeIdList = append(nodeIdList, v.NodeId)
	}
	if len(nodeIdList) > 0 {
		db.Where("status=1").Find(&nodes, nodeIdList)
	}

	var nodeServices []models.NodeService
	var serviceIdList []uint32
	var services []models.Service
	db.Where("node_id IN ?", nodeIdList).Find(&nodeServices)
	for _, v := range nodeServices {
		serviceIdList = append(serviceIdList, v.ServiceId)
	}
	if len(serviceIdList) > 0 {
		serviceIdList = Unique(serviceIdList)
		db.Find(&services, serviceIdList)
	}

	reqNodes := []reqNode{}
	for _, v := range nodes {
		serverNameList := []string{}
		for _, vv := range nodeServices {
			if vv.NodeId == v.Id {
				for _, vvv := range services {
					if vvv.Id == vv.ServiceId {
						serverNameList = append(serverNameList, vvv.Sign)
					}
				}
			}
		}
		serverNameList = UniqueString(serverNameList)
		reqNodes = append(reqNodes, reqNode{
			Addr:    v.Addr,
			Service: strings.Join(serverNameList, ","),
			Weight:  int(v.Weight),
		})
	}

	crashList := []string{}
	lower := map[string]int{}
	for _, v := range services {
		if v.Status == 0 {
			crashList = append(crashList, v.Sign)
		}
		if v.Status > 0 && v.Status < 100 {
			lower[v.Sign] = int(v.Status)
		}
	}

	reqData := reqSync{
		Cmd: "sync",
		Data: reqSyncData{
			Nodes: reqNodes,
			Crash: strings.Join(crashList, ","),
			Lower: lower,
		},
	}

	apiRequest(gw.Api, gw.Key, reqData)
}
