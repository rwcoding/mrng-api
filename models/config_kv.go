package models

type ConfigKv struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	K         string
	V         string
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *ConfigKv) TableName() string {
	return "config_kv"
}
