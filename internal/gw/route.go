package gw

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB

func init() {
	goback.Route("mrng.gw.list", NewApiList, "网关列表")
	goback.Route("mrng.gw.add", NewApiAdd, "网关增加")
	goback.Route("mrng.gw.edit", NewApiEdit, "网关编辑")
	goback.Route("mrng.gw.info", NewApiInfo, "网关详情")
	goback.Route("mrng.gw.delete", NewApiDelete, "网关删除")
	goback.Route("mrng.gw.status", NewApiStatus, "网关状态")
	goback.Route("mrng.gw.bind.query", NewApiBindQuery, "绑定节点查询")
	goback.Route("mrng.gw.bind", NewApiBind, "绑定节点")
	goback.Route("mrng.gw.sync", NewApiSync, "同步网关")
}
