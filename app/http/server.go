package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RudyChow/code-snippets/app/redis"
	"github.com/RudyChow/code-snippets/conf"
)

// StartServer : 开启http服务
func StartServer() {
	r := gin.Default()
	api := r.Group("/api")
	api.GET("/snippet/:id", getSnippet)
	api.POST("/snippet", storeSnippet)
	r.Run(conf.Cfg.Http.Addr) // listen and serve on
}

// 获取片段
func getSnippet(c *gin.Context) {
	id := c.Param("id")
	snippet, err := redis.RedisClient.GetSnippet(id)

	c.JSON(http.StatusOK, getResponse(snippet, err))
}

// 保存片段
func storeSnippet(c *gin.Context) {
	var snippet *redis.Snippet

	if err := c.ShouldBind(&snippet); err != nil {
		c.JSON(http.StatusOK, getResponse(nil, err))
		return
	}

}

func getResponse(data interface{}, err error) map[string]interface{} {
	response := make(map[string]interface{})

	if err == nil {
		response["error"] = ""
	} else {
		log.Println(err)
		response["error"] = "something went wrong"
	}
	response["data"] = data

	return response
}
