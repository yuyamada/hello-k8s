package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuyamada/hello-k8s/app/article"
	"github.com/yuyamada/hello-k8s/app/httpd/handler"
)

func main() {
	articles := article.New()
	r := gin.Default()
	r.GET("/article", handler.ArticlesGet(articles))
	r.POST("/article", handler.ArticlesPost(articles))

	r.Run()
}
