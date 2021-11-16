package models

type ConfigWhite struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Ip        string
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *ConfigWhite) TableName() string {
	return "config_white"
}
