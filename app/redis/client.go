package redis

import (
	"log"

	"github.com/RudyChow/code-snippets/conf"
	"github.com/go-redis/redis"
)

// RedisClient redis客户端
var RedisClient *redisDriver

type redisDriver struct {
	client *redis.Client
}

// RedisSnippetID 自增id
// RedisSnippetDetail 片段详情
// RedisSnippetExpire 过期时间
const (
	RedisSnippetID     = "snippets:id"
	RedisSnippetDetail = "snippets:detail:"
	RedisSnippetExpire = 24 * 60 * 60
)

//新建一个redis客户端
func init() {

	client := redis.NewClient(&redis.Options{
		Addr:     conf.Cfg.Redis.Addr,
		Password: conf.Cfg.Redis.Auth, // no password set
		DB:       conf.Cfg.Redis.DB,   // use default DB
	})

	if _, err := client.Ping().Result(); err != nil {
		log.Panic("redis connect failed")
	}
	RedisClient = &redisDriver{
		client: client,
	}
}
