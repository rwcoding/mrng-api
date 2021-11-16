package models

type ConfigMain struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Sign      string
	Env       string
	Project   string
	K         string
	V         string
	Status    uint8
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *ConfigMain) TableName() string {
	return "config_main"
}
