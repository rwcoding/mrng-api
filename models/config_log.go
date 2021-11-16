package models

const LOG_TYPE_CREATE = 1
const LOG_TYPE_UPDATE = 2
const LOG_TYPE_DELETE = 3

func LogTypesNames() map[int]string {
	return map[int]string{
		LOG_TYPE_CREATE: "创建",
		LOG_TYPE_UPDATE: "更新",
		LOG_TYPE_DELETE: "删除",
	}
}

type ConfigLog struct {
	Id        uint32 `gorm:"primary_key;AUTO_INCREMENT"`
	Type      uint8
	Name      string
	Sign      string
	Env       string
	K         string
	V         string
	Project   string
	AdminerId uint32
	CreatedAt uint32
	UpdatedAt uint32
}

func (m *ConfigLog) TableName() string {
	return "config_log"
}
