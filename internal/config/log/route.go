package log

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB

func init() {
	goback.Route("mrng.config.log.list", NewApiList, "日志列表")
	goback.Route("mrng.config.log.info", NewApiInfo, "日志详情")
}
