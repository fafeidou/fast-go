package routers

import (
	"fast-go/internal/app/controller/sys"
	"github.com/gin-gonic/gin"
)

func RegisterRouterSys(app *gin.RouterGroup) {

	article := sys.Article{}
	app.GET("/article/list", article.List)
	app.POST("/article/update", article.EditArticle)

	//app.GET("/menu/detail", menu.Detail)
	//app.GET("/menu/allmenu", menu.AllMenu)
	//app.GET("/menu/menubuttonlist", menu.MenuButtonList)
	//app.POST("/menu/delete", menu.Delete)
	//app.POST("/menu/update", menu.Update)
	//app.POST("/menu/create", menu.Create)
}
