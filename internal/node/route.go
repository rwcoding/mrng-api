package node

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB

func init() {
	goback.Route("mrng.node.list", NewApiList, "节点列表")
	goback.Route("mrng.node.add", NewApiAdd, "节点增加")
	goback.Route("mrng.node.edit", NewApiEdit, "节点编辑")
	goback.Route("mrng.node.info", NewApiInfo, "节点详情")
	goback.Route("mrng.node.delete", NewApiDelete, "节点删除")
	goback.Route("mrng.node.status", NewApiStatus, "节点上下架")
	goback.Route("mrng.node.bind.query", NewApiBindQuery, "节点服务绑定查询")
	goback.Route("mrng.node.bind", NewApiBind, "节点服务绑定")
}
