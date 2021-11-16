package models

import (
	"github.com/rwcoding/mrng/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func init() {
	InitDB()
}

func InitDB() *gorm.DB {
	lvl := logger.Warn
	if config.IsDev() {
		lvl = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), //需要文件的替换为文件
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  lvl,         // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	conf := config.GetDb()
	if conf.Username == "" && conf.Password == "" {
		return nil
	}
	dsn := conf.Username + ":" + conf.Password + "@tcp(" +
		conf.Host + ":" + strconv.Itoa(conf.Port) + ")/" +
		conf.Dbname + "?charset=" + conf.Charset + "&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal(err)
	}

	//连接池
	sd, _ := db.DB()
	if conf.MaxOpenConn == 0 {
		conf.MaxOpenConn = 100
	}
	if conf.MaxIdleConn == 0 {
		conf.MaxIdleConn = 50
	}
	if conf.MaxLifetime == 0 {
		conf.MaxLifetime = 600
	}
	sd.SetMaxOpenConns(conf.MaxOpenConn)
	sd.SetMaxIdleConns(conf.MaxIdleConn)
	sd.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifetime))

	return db
}
