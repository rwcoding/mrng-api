package models

import "gorm.io/plugin/soft_delete"

type Node struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Addr      string
	Services  string
	Status    uint8
	Weight    uint8
	CreatedAt uint32
	UpdatedAt uint32
	DeletedAt soft_delete.DeletedAt
}

func (m *Node) TableName() string {
	return "mrng_node"
}
