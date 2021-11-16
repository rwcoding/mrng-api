package kv

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB()

func init() {
	goback.Route("mrng.config.kv.list", NewApiList, "缓存列表")
	goback.Route("mrng.config.kv.info", NewApiInfo, "缓存详情")
}
