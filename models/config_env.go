package models

type ConfigEnv struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string
	Sign      string
	Ord       uint32
	KeyV1     string
	KeyV2     string
	Version   int64
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *ConfigEnv) TableName() string {
	return "config_env"
}
