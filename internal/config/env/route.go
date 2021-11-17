package env

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB

func init() {
	goback.Route("mrng.config.env.list", NewApiList, "环境列表")
	goback.Route("mrng.config.env.add", NewApiAdd, "环境增加")
	goback.Route("mrng.config.env.edit", NewApiEdit, "环境编辑")
	goback.Route("mrng.config.env.info", NewApiInfo, "环境详情")
	goback.Route("mrng.config.env.delete", NewApiDelete, "环境删除")
}
