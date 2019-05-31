package redis

import (
	"encoding/json"
	"time"

	"github.com/RudyChow/code-snippets/app/utils"
	"github.com/RudyChow/code-snippets/conf"
)

// IncrID : 自增id
func (redis *redisDriver) IncrID() int64 {
	id := redis.client.Incr(conf.Cfg.Redis.Snippet.IncrKey).Val()

	return id
}

// StoreSnippet : 保存代码片段
func (redis *redisDriver) StoreSnippet(id string, value *Snippet) error {
	_, err := redis.client.Set(conf.Cfg.Redis.Snippet.DetailKey+id, value, time.Duration(conf.Cfg.Redis.Snippet.Expire)*time.Second).Result()

	return err
}

// GetSnippet : 获取代码片段
func (redis *redisDriver) GetSnippet(id string) (*Snippet, error) {
	var snippet *Snippet

	result, err := redis.client.Get(conf.Cfg.Redis.Snippet.DetailKey + id).Result()
	if err == nil {
		json.Unmarshal([]byte(result), &snippet)
	}
	return snippet, err
}

// AutoStoreSnippet : 自动保存代码片段
func (redis *redisDriver) AutoStoreSnippet(value *Snippet) (string, error) {
	id := redis.IncrID()
	short := utils.GenerateShort(id)
	_, err := redis.client.Set(conf.Cfg.Redis.Snippet.DetailKey+short, value, RedisSnippetExpire*time.Second).Result()
	if err != nil {
		short = ""
	}

	return short, err
}
