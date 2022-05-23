package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(app *gin.Engine, baseUrl string) {
	g := app.Group(baseUrl)
	{
		userGroup := g.Group("/user")
		{
			userGroup.POST("/", warp(userRegister))
		}

		urlGroup := g.Group("/url")
		{
			urlGroup.POST("", warp(addUrl))
			urlGroup.GET("/:id", warp(getUrl))
			urlGroup.GET("/:id/go", warp(visitUrl))
			urlGroup.GET("/:id/visit/:visitId", warp(getUrlVisit))
			urlGroup.GET("/:id/visit", warp(urlVisitList))
		}
	}
}
