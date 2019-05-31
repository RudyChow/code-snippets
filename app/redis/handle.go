package redis

import (
	"encoding/json"
	"time"
)

// IncrID : 自增id
func (redis *redisDriver) IncrID() int64 {
	id := redis.client.Incr(RedisSnippetID).Val()

	return id
}

// StoreSnippet : 保存代码片段
func (redis *redisDriver) StoreSnippet(id string, value *Snippet) error {
	_, err := redis.client.Set(RedisSnippetDetail+id, value, RedisSnippetExpire*time.Second).Result()

	return err
}

// GetSnippet : 获取代码片段
func (redis *redisDriver) GetSnippet(id string) (*Snippet, error) {
	var snippet *Snippet

	result, err := redis.client.Get(RedisSnippetDetail + id).Result()
	if err == nil {
		json.Unmarshal([]byte(result), &snippet)
	}
	return snippet, err
}
