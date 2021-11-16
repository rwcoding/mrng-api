package service

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB()

func init() {
	goback.Route("mrng.service.list", NewApiList, "服务列表")
	goback.Route("mrng.service.add", NewApiAdd, "服务增加")
	goback.Route("mrng.service.edit", NewApiEdit, "服务编辑")
	goback.Route("mrng.service.info", NewApiInfo, "服务详情")
	goback.Route("mrng.service.delete", NewApiDelete, "服务删除")
	goback.Route("mrng.service.recover", NewApiRecover, "服务恢复")
}
