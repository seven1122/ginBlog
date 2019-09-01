package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/example/basic/docs"

	"seven1122/ginBlog/middleware/jwt"
	"seven1122/ginBlog/routers/api"
	"seven1122/ginBlog/routers/api/v1"

	"seven1122/ginBlog/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)
	apiV1 := r.Group("api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.EditArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.POST("/auth", api.GetAuth)

	return r
}
