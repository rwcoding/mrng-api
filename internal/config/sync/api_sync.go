package sync

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
)

type syncRequest struct {
	ctx *boot.Context
}

func NewApiSync(ctx *boot.Context) boot.Logic {
	return &syncRequest{ctx: ctx}
}

func (request *syncRequest) Run() *api.Response {

	var keys []string

	//环境
	var envs []models.ConfigEnv
	db().Find(&envs)
	for _, v := range envs {
		keys = append(keys, services.CacheKeyForEnv(v.Sign, "v1"))
		keys = append(keys, services.CacheKeyForEnv(v.Sign, "v2"))
		_ = services.SetCacheForEnv(v.Sign, v.KeyV1, v.KeyV2)
	}

	//工程
	var projects []models.ConfigProject
	db().Find(&projects)
	for _, v := range projects {
		keys = append(keys, services.CacheKeyForProject(v.Sign, "v1"))
		keys = append(keys, services.CacheKeyForProject(v.Sign, "v2"))
		_ = services.SetCacheForProject(v.Sign, v.KeyV1, v.KeyV2)
	}

	//白名单
	var ips []models.ConfigWhite
	db().Find(&ips)
	for _, v := range ips {
		keys = append(keys, services.CacheKeyForWhite(v.Ip))
		_ = services.SetCacheForWhite(v.Ip)
	}

	//配置
	var configs []models.ConfigMain
	db().Find(&configs)
	for _, v := range configs {
		keys = append(keys, services.CacheKeyForConfig(v.Sign))
		_ = services.SetCacheForConfig(v.Sign, v.V)
	}

	keys = append(keys, services.CacheKeyForVersion())
	_ = services.SetVersion("")

	//清除多余陈旧的缓存
	var kvs []models.ConfigKv
	db().Select("k").Find(&kvs)
	for _, v := range kvs {
		have := false
		for _, vv := range keys {
			if vv == v.K {
				have = true
			}
		}
		if !have {
			_ = services.DeleteCache(v.K)
		}
	}

	return api.NewSuccessResponse("同步成功")
}
