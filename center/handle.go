package center

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/rwcoding/mrng/models"
	"github.com/rwcoding/mrng/services"
	"io/ioutil"
	"net/http"
	"strings"
)

var db = models.GetDB()

func failure(msg string) []byte {
	res := response{
		Code:     0,
		Msg:      msg,
		Versions: map[string]int64{},
		Data:     map[string]interface{}{},
	}
	b, _ := json.Marshal(res)
	return b
}

func success(versions map[string]int64, data map[string]interface{}) ([]byte, error) {
	res := response{
		Code:     1,
		Msg:      "",
		Data:     data,
		Versions: versions,
	}
	b, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type request struct {
	Certs map[string]string `json:"certs"`
	Keys  []string          `json:"keys"`
}

type response struct {
	Code     int                    `json:"code"`
	Msg      string                 `json:"msg"`
	Versions map[string]int64       `json:"versions"`
	Data     map[string]interface{} `json:"data"`
}

func Handle(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Data(http.StatusOK, "application/json", failure("数据读取错误"))
		return
	}

	req := &request{}
	if json.Unmarshal(body, req) != nil {
		c.Data(http.StatusOK, "application/json", failure("json解析错误"))
		return
	}

	certs := map[string]string{}
	needVerifyEnv := []string{}
	needVerifyProject := []string{}

	for _, v := range req.Keys {
		tmp := strings.Split(v, ".")
		if len(tmp) < 3 {
			continue
		}
		needVerifyEnv = append(needVerifyEnv, tmp[0])
		needVerifyProject = append(needVerifyProject, tmp[1])
	}
	for k, _ := range req.Certs {
		tmp := strings.Split(k, ":")
		if len(tmp) < 2 {
			c.Data(http.StatusOK, "application/json", failure(k+"错误"))
			return
		}
		if tmp[0] == "E" {
			needVerifyEnv = append(needVerifyEnv, tmp[1])
		}
		if tmp[0] == "P" {
			needVerifyProject = append(needVerifyProject, tmp[1])
		}
	}

	ret := map[string]interface{}{}

	//验证密钥
	for _, env := range needVerifyEnv {
		ck1 := "E:" + env + ":v1"
		ck2 := "E:" + env + ":v2"
		_, ok := certs[ck1]
		if !ok {
			certs[ck1] = getKey(services.CacheKeyForEnv(env, "v1")).String()
			certs[ck2] = getKey(services.CacheKeyForEnv(env, "v2")).String()
		}

		if certs[ck1] != "" || certs[ck2] != "" {
			if req.Certs["E:"+env] != certs[ck1] && req.Certs["E:"+env] != certs[ck2] {
				c.Data(http.StatusOK, "application/json", failure("环境"+env+"密钥错误"))
				return
			}
		}
	}

	for _, project := range needVerifyProject {
		k1 := "P:" + project + ":v1"
		k2 := "P:" + project + ":v2"
		_, ok := certs[k1]
		if !ok {
			certs[k1] = getKey(services.CacheKeyForProject(project, "v1")).String()
			certs[k2] = getKey(services.CacheKeyForProject(project, "v2")).String()
		}

		if certs[k1] != "" || certs[k2] != "" {
			if req.Certs["P:"+project] != certs[k1] && req.Certs["P:"+project] != certs[k2] {
				c.Data(http.StatusOK, "application/json", failure("工程"+project+"密钥错误"))
				return
			}
		}
	}

	for _, v := range req.Keys {
		if vv, err := getKey(services.CacheKeyForConfig(v)).Result(); err == nil {
			ret[v] = vv
		} else {
			ret[v] = nil
		}
	}

	versions := map[string]int64{}
	if vs := getKey(services.CacheKeyForVersion()).String(); vs != "" {
		_ = json.Unmarshal([]byte(vs), &versions)
	}

	res, err := success(versions, ret)
	if err != nil {
		c.String(http.StatusOK, "application/json", failure(err.Error()))
		return
	}

	c.Data(http.StatusOK, "application/json", res)
}

func getKey(key string) *cacheResult {
	rc := services.NewRedis()
	if rc == nil {
		var kv models.ConfigKv
		db.Where("k=?", key).Take(&kv)
		return &cacheResult{
			ds: &kv,
		}
	}
	return &cacheResult{
		rs: rc.Get(key),
	}
}

type cacheResult struct {
	rs *redis.StringCmd
	ds *models.ConfigKv
}

func (rr *cacheResult) String() string {
	if rr.rs == nil {
		if rr.ds != nil {
			return rr.ds.V
		}
		return ""
	}
	s, err := rr.rs.Result()
	if err != nil {
		return ""
	}
	return s
}

func (rr *cacheResult) Result() (string, error) {
	if rr.rs == nil {
		if rr.ds != nil && rr.ds.Id > 0 {
			return rr.ds.V, nil
		}
		return "", nil
	}
	return rr.rs.Result()
}
