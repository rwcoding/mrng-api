package models

import "gorm.io/plugin/soft_delete"

type GwNode struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	GwId      uint32
	NodeId    uint32
	CreatedAt uint32
	UpdatedAt uint32
	DeletedAt soft_delete.DeletedAt
}

func (m *GwNode) TableName() string {
	return "mrng_gw_node"
}
