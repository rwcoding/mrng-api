package models

import "gorm.io/plugin/soft_delete"

type Service struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Sign      string
	Status    uint8
	CreatedAt uint32
	UpdatedAt uint32
	DeletedAt soft_delete.DeletedAt
}

func (m *Service) TableName() string {
	return "mrng_service"
}
