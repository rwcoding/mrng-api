package main1

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB()

func init() {
	goback.Route("mrng.config.main.list", NewApiList, "配置列表")
	goback.Route("mrng.config.main.add", NewApiAdd, "配置增加")
	goback.Route("mrng.config.main.edit", NewApiEdit, "配置编辑")
	goback.Route("mrng.config.main.info", NewApiInfo, "配置详情")
	goback.Route("mrng.config.main.delete", NewApiDelete, "配置删除")
	goback.Route("mrng.config.main.status", NewApiStatus, "配置发布")
}
