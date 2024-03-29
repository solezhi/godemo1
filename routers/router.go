package routers

import (
	"demo2/pkg/setting"
	"demo2/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	apiv1 := r.Group("/api/v1")

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	apiv1.GET("/tags", v1.GetTags)
	apiv1.POST("/tags", v1.AddTag)
	apiv1.PUT("/tags/:id", v1.EditTag)
	apiv1.DELETE("/tags/:id", v1.DeleteTag)

	apiv1.GET("/articles", v1.GetArticles)
	apiv1.GET("/articles/:id", v1.GetArticle)
	apiv1.POST("/articles", v1.AddArticle)
	apiv1.PUT("/articles/:id", v1.EditArticle)
	apiv1.DELETE("/articles/:id", v1.DeleteArticle)

	return r
}
