package sync

import (
	"github.com/rwcoding/goback"
	"github.com/rwcoding/mrng/models"
)

var db = models.GetDB

func init() {
	goback.Route("mrng.config.sync", NewApiSync, "同步到缓存")
}
