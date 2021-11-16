package models

type ConfigProject struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Sign      string
	Ord       uint32
	KeyV1     string
	KeyV2     string
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *ConfigProject) TableName() string {
	return "config_project"
}
