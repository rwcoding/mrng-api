package services

import (
	"encoding/json"
	"errors"
	"github.com/rwcoding/mrng/models"
	"strings"
)

func DeleteCache(k string) error {
	var kv models.ConfigKv
	db().Where("k=?", k).Delete(&kv)
	for _, v := range redisPools {
		v.Del(k)
	}
	return nil
}

func SetCache(k, v string) error {
	var kv models.ConfigKv
	db().Where("k=?", k).Take(&kv)
	if kv.Id > 0 {
		if kv.V != v {
			tx := db().Model(kv).Update("v", v)
			if tx.Error != nil {
				return tx.Error
			}
		}
	} else {
		kv.K = k
		kv.V = v
		tx := db().Create(&kv)
		if tx.Error != nil {
			return tx.Error
		}
	}

	for _, r := range redisPools {
		r.Set(k, v)
	}

	return nil
}

func GetCache(k string) (string, error) {
	var kv models.ConfigKv
	db().Where("k=?", k).Take(&kv)
	if kv.Id > 0 {
		return kv.V, nil
	}
	return "", errors.New("not found")
}

func SetCacheForEnv(sign, v1, v2 string) error {
	_ = SetCache(CacheKeyForEnv(sign, "v1"), v1)
	return SetCache(CacheKeyForEnv(sign, "v2"), v2)
}

func DeleteCacheForEnv(sign string) error {
	_ = DeleteCache(CacheKeyForEnv(sign, "v1"))
	return DeleteCache(CacheKeyForEnv(sign, "v2"))
}

func CacheKeyForEnv(sign, ver string) string {
	return "E:" + sign + ":" + ver
}

func SetCacheForProject(sign, v1, v2 string) error {
	_ = SetCache(CacheKeyForProject(sign, "v1"), v1)
	return SetCache(CacheKeyForProject(sign, "v2"), v2)
}

func DeleteCacheForProject(sign string) error {
	_ = DeleteCache(CacheKeyForProject(sign, "v1"))
	return DeleteCache(CacheKeyForProject(sign, "v2"))
}

func CacheKeyForProject(sign, ver string) string {
	return "P:" + sign + ":" + ver
}

func SetCacheForWhite(ip string) error {
	return SetCache(CacheKeyForWhite(ip), "1")
}

func DeleteCacheForWhite(ip string) error {
	return DeleteCache(CacheKeyForWhite(ip))
}

func CacheKeyForWhite(ip string) string {
	return "W:" + ip
}

func SetCacheForConfig(sign, v string) error {
	_ = SetCache(CacheKeyForConfig(sign), v)
	return SetVersion(sign)
}

func DeleteCacheForConfig(sign string) error {
	_ = DeleteCache(CacheKeyForConfig(sign))
	return SetVersion(sign)
}

func CacheKeyForConfig(sign string) string {
	return "C:" + sign
}

func CacheKeyForVersion() string {
	return "VERSIONS"
}

func SetVersion(sign string) error {
	if sign != "" {
		arr := strings.Split(sign, ".")
		db().Exec("UPDATE "+(&models.ConfigEnv{}).TableName()+" SET version = version+1 WHERE sign = ?", arr[0])
	}

	var envs []models.ConfigEnv
	versions := map[string]interface{}{}
	db().Find(&envs)

	if len(envs) == 0 {
		return nil
	}

	for _, v := range envs {
		versions[v.Sign] = v.Version
	}
	b, err := json.Marshal(versions)
	if err != nil {
		return err
	}
	return SetCache(CacheKeyForVersion(), string(b))
}
