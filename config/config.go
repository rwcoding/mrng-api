package config

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/rwcoding/goback"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var _config *config

type RedisConfig struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	Pool     int    `toml:"pool"`
}

type DbConfig struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	Username    string `toml:"username"`
	Password    string `toml:"password"`
	Dbname      string `toml:"dbname"`
	Charset     string `toml:"charset"`
	MaxOpenConn int    `toml:"pool_max_open"`
	MaxIdleConn int    `toml:"pool_max_idle"`
	MaxLifetime int    `toml:"pool_max_life"`
}

type config struct {
	Lang   string        `toml:"lang"`
	Addr   string        `toml:"addr"`
	Env    string        `toml:"env"`
	Log    string        `toml:"log"`
	OnlyGP bool          `toml:"only_gp"`
	Header bool          `toml:"header"`
	Db     DbConfig      `toml:"db"`
	Timer  int           `toml:"timer"`
	OnlyCc bool          `toml:"only_cc"`
	Redis  []RedisConfig `toml:"redis"`
}

func init() {
	c := config{
		Timer:  3600,
		OnlyCc: false,
		Lang:   "zh",
		Addr:   ":8080",
		Env:    "prod",
		OnlyGP: false,
		Header: false,
	}

	b, err := ReadConfigFile(goback.App().Console.GetFlag("conf"))
	_, err = toml.Decode(string(b), &c)
	if err != nil {
		log.Fatal("failed to parse config file, ", err.Error())
	}
	_config = &c
}

func GetAddr() string {
	return _config.Addr
}

func GetTimer() int {
	return _config.Timer
}

func IsOnlyCc() bool {
	return _config.OnlyCc
}

func IsOnlyGP() bool {
	return _config.OnlyGP
}

func IsWriteHeader() bool {
	return _config.Header
}

func GetRedis() []RedisConfig {
	return _config.Redis
}

func GetLang() string {
	return _config.Lang
}

func GetEnv() string {
	return _config.Env
}

func GetLog() string {
	if _config.Log == "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		_config.Log = dir + string(os.PathSeparator) + "mrng.{ymd}.log"
	}
	return _config.Log
}

func GetMode() string {
	if _config.Env == "dev" {
		return gin.DebugMode
	}
	return gin.ReleaseMode
}

func GetDb() DbConfig {
	return _config.Db
}

func IsDev() bool {
	return _config.Env == "dev"
}

func ReadConfigFile(configFile string) (bytes []byte, err error) {
	var f http.File
	var dir string

	if configFile != "" {
		f, err = os.Open(configFile)
	} else {
		if dir, err = filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
			f, err = os.Open(dir + string(os.PathSeparator) + "mrng.toml")
			if err != nil {
				if dir, err = os.Getwd(); err == nil {
					f, err = os.Open(dir + string(os.PathSeparator) + "mrng.toml")
				}
			}
		}
	}
	if err != nil {
		log.Fatal("failed to parse config file, ", err.Error())
		return
	}

	defer f.Close()

	return ioutil.ReadAll(f)
}
