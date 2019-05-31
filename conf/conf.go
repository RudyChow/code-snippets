package conf

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Cfg : 配置
var Cfg *config

type config struct {
	Redis *redis
	HTTP  *http
}

type redis struct {
	Addr    string
	Auth    string
	DB      int
	Snippet *snippet
}

type http struct {
	Addr string
	Mode string
}

type snippet struct {
	IncrKey   string
	DetailKey string
	Expire    int
}

func init() {
	//读取配置文件
	if _, err := toml.DecodeFile("./config.toml", &Cfg); err != nil {
		log.Panic(err)
	}
}
