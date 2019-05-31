package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

var Cfg *Config

type Config struct {
	Redis *redis
	Http  *http
}

type redis struct {
	Addr string
	Auth string
	DB   int
}

type http struct {
	Addr string
}

func init() {
	//读取配置文件
	if _, err := toml.DecodeFile("./config.toml", &Cfg); err != nil {
		log.Panic(err)
	}
}
