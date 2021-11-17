package white

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"strings"
)

type syncNodeRequest struct {
	ctx *boot.Context
}

func NewApiSyncNode(ctx *boot.Context) boot.Logic {
	return &syncNodeRequest{ctx: ctx}
}

func (request *syncNodeRequest) Run() *api.Response {

	var ips []string
	var nodes []models.Node
	db().Find(&nodes)
	for _, v := range nodes {
		tmp := strings.Split(v.Addr, ":")
		if len(tmp) > 0 && services.VerifyIp(tmp[0]) {
			ips = append(ips, tmp[0])
		}
	}

	if len(ips) == 0 {
		return api.NewErrorResponse("同步失败")
	}
	ips = services.UniqueString(ips)

	var ws []models.ConfigWhite
	db().Where("ip IN ?", ips).Find(&ws)

	var insert []models.ConfigWhite
	for _, v := range ips {
		have := false
		for _, vv := range ws {
			if vv.Ip == v {
				have = true
			}
		}
		if !have {
			insert = append(insert, models.ConfigWhite{Ip: v})
		}
	}

	if len(insert) > 0 {
		db().Create(insert)
		for _, v := range insert {
			_ = services.SetCacheForWhite(v.Ip)
		}
	}

	return api.NewSuccessResponse("同步成功")
}
