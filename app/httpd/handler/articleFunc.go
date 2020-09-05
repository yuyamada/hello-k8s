package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuyamada/hello-k8s/app/article"
)

// ArticlesGet is handler func for Articles with GET method
func ArticlesGet(articles *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := articles.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

// ArticlesPostRequest is request body for Articles with POST method
type ArticlesPostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// ArticlesPost is handler func for Articles with POST method
func ArticlesPost(post *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := ArticlesPostRequest{}
		c.Bind(&requestBody)

		item := article.Item{
			Title:       requestBody.Title,
			Description: requestBody.Description,
		}
		post.Add(item)

		c.Status(http.StatusNoContent)
	}
}
