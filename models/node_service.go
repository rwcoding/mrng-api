package models

import "gorm.io/plugin/soft_delete"

type NodeService struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	ServiceId uint32
	NodeId    uint32
	CreatedAt uint32
	UpdatedAt uint32
	DeletedAt soft_delete.DeletedAt
}

func (m *NodeService) TableName() string {
	return "mrng_node_service"
}
