package services

import (
	"github.com/rwcoding/mrng/models"
)

type reqDelete struct {
	Addr string `json:"addr"`
}

type reqSyncDelete struct {
	Cmd  string    `json:"cmd"`
	Data reqDelete `json:"data"`
}

func SyncNodeDelete(node models.Node) {
	reqData := reqSyncDelete{
		Cmd: "delete",
		Data: reqDelete{
			Addr: node.Addr,
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
