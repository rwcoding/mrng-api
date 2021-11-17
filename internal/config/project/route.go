package project

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB

func init() {
	goback.Route("mrng.config.project.list", NewApiList, "工程列表")
	goback.Route("mrng.config.project.add", NewApiAdd, "工程增加")
	goback.Route("mrng.config.project.edit", NewApiEdit, "工程编辑")
	goback.Route("mrng.config.project.info", NewApiInfo, "工程详情")
	goback.Route("mrng.config.project.delete", NewApiDelete, "工程删除")
}
