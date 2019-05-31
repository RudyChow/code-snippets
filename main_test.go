package main

import (
	"testing"

	"github.com/RudyChow/code-snippets/app/redis"
	"github.com/RudyChow/code-snippets/app/utils"
)

func TestMain(t *testing.T) {
	//自增
	id := redis.RedisClient.IncrID()
	short := utils.GenerateShort(id)

	snippet := &redis.Snippet{
		Language: "c++",
		Version:  "1.66",
		Code:     "<?php\necho 1;",
	}
	//保存代码片段
	err := redis.RedisClient.StoreSnippet(short, snippet)

	if err != nil {
		t.Error(err)
	}

	//获取代码片段
	result, err := redis.RedisClient.GetSnippet(short)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)

	//获取ttl
	ttl, err := redis.RedisClient.GetSnippetTTL(short)
	if err != nil {
		t.Error(err)
	}

	t.Log(ttl)

	//重新设置过期时间
	err = redis.RedisClient.ExpireSnippet(short)
	if err != nil {
		t.Error(err)
	}

}
