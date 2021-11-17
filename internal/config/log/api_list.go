package log

import (
	"github.com/rwcoding/goback/pkg/api"
	"github.com/rwcoding/goback/pkg/boot"
	"github.com/rwcoding/mrng/models"
)

type listRequest struct {
	ctx *boot.Context

	Page     int    `validate:"required,numeric,min=1" json:"page"`
	PageSize int    `validate:"omitempty,numeric,max=20" json:"page_size"`
	Name     string `validate:"omitempty,max=100" json:"name"`
	K        string `validate:"omitempty,max=100" json:"k"`
}

type itemResponse struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Sign    string `json:"sign"`
	Env     string `json:"env"`
	Project string `json:"project"`
	K       string `json:"k"`
	V       string `json:"v"`
	Type    uint8  `json:"type"`
}

type listResponse struct {
	Datas        []itemResponse    `json:"datas"`
	Count        int64             `json:"count"`
	Page         int               `json:"page"`
	PageSize     int               `json:"page_size"`
	TypeNames    map[int]string    `json:"type_names"`
	EnvNames     map[string]string `json:"env_names"`
	ProjectNames map[string]string `json:"project_names"`
}

func NewApiList(ctx *boot.Context) boot.Logic {
	return &listRequest{ctx: ctx}
}

func (request *listRequest) Run() *api.Response {
	pageSize := request.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (request.Page - 1) * pageSize
	var us []models.ConfigLog
	var c int64

	tx1 := db().Order("id desc").Offset(offset).Limit(pageSize)
	tx2 := db().Model(&models.ConfigLog{})
	if request.Name != "" {
		tx1.Where("name LIKE ?", "%"+request.Name+"%")
		tx2.Where("name LIKE ?", "%"+request.Name+"%")
	}
	if request.K != "" {
		tx1.Where("k LIKE ?", "%"+request.K+"%")
		tx2.Where("k LIKE ?", "%"+request.K+"%")
	}
	tx1.Find(&us)
	tx2.Count(&c)

	var list []itemResponse
	for _, v := range us {
		list = append(list, itemResponse{
			Id:      v.Id,
			Name:    v.Name,
			Sign:    v.Sign,
			Env:     v.Env,
			Project: v.Project,
			K:       v.K,
			V:       v.V,
			Type:    v.Type,
		})
	}

	var es []models.ConfigEnv
	var ps []models.ConfigProject
	db().Order("ord").Find(&es)
	db().Order("ord").Find(&ps)

	eNames := map[string]string{}
	pNames := map[string]string{}
	for _, v := range es {
		eNames[v.Sign] = v.Name
	}
	for _, v := range ps {
		pNames[v.Sign] = v.Name
	}

	return api.NewDataResponse(&listResponse{
		Datas:        list,
		Count:        c,
		Page:         request.Page,
		PageSize:     pageSize,
		TypeNames:    models.LogTypesNames(),
		EnvNames:     eNames,
		ProjectNames: pNames,
	})
}
