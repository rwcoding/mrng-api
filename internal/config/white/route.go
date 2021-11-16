package white

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB()

func init() {
	goback.Route("mrng.config.white.list", NewApiList, "白名单列表")
	goback.Route("mrng.config.white.add", NewApiDelete, "白名单增加")
	goback.Route("mrng.config.white.edit", NewApiEdit, "白名单编辑")
	goback.Route("mrng.config.white.info", NewApiInfo, "白名单详情")
	goback.Route("mrng.config.white.delete", NewApiDelete, "白名单删除")
	goback.Route("mrng.config.white.node", NewApiSyncNode, "白名单同步节点")
}
