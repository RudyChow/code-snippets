package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RudyChow/code-snippets/app/redis"
	"github.com/RudyChow/code-snippets/conf"
)

// StartServer : 开启http服务
func StartServer() {
	gin.SetMode(conf.Cfg.HTTP.Mode)
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/snippet/:id", getSnippet)
	api.POST("/snippet", storeSnippet)
	r.Run(conf.Cfg.HTTP.Addr) // listen and serve on
}

// 获取片段
func getSnippet(c *gin.Context) {
	id := c.Param("id")
	snippet, err := redis.RedisClient.GetSnippet(id)

	c.JSON(http.StatusOK, getResponse(snippet, err))
}

// 保存片段
func storeSnippet(c *gin.Context) {
	var snippet redis.Snippet

	if err := c.ShouldBind(&snippet); err != nil {
		c.JSON(http.StatusOK, getResponse(nil, err))
		return
	}

	short, err := redis.RedisClient.AutoStoreSnippet(&snippet)

	c.JSON(http.StatusOK, getResponse(short, err))
}

// 生成响应
func getResponse(data interface{}, err error) map[string]interface{} {
	response := make(map[string]interface{})

	if err == nil {
		response["error"] = ""
	} else {
		response["error"] = err.Error()
	}
	response["data"] = data

	return response
}
