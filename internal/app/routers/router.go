package routers

import (
	"fast-go/internal/app/controller/common"
	"fast-go/internal/app/controller/sys"
	"fast-go/internal/app/models"
	"fast-go/middleware/jwt"
	jwtauth "fast-go/middleware/jwtAuth"
	"fmt"
	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/auth", sys.GetAuth)
	r.POST("/upload", sys.UploadImage)
	var articles []models.Article
	maps := make(map[string]interface{})
	aa := models.Find(&articles, 1, 10, maps)
	r.GET("/", func(c *gin.Context) { common.ResSuccessPage(c, 20, aa) })
	apiPrefix := "/api"
	g := r.Group(apiPrefix)
	g.Use(jwt.JWT())
	{
		article := sys.Article{}
		g.GET("/article/list", article.List)
		g.POST("/article/update", article.EditArticle)
		//RegisterRouterSys(g)
	}

	r.GET("/jwt", func(c *gin.Context) {
		j := &jwtauth.JWT{
			[]byte("test"),
		}
		claims := jwtauth.CustomClaims{
			1,
			"awh521",
			"1044176017@qq.com",

			jwt2.StandardClaims{
				ExpiresAt: 15000, //time.Now().Add(24 * time.Hour).Unix()
				Issuer:    "test",
			},
		}
		token, err := j.CreateToken(claims)
		if err != nil {
			c.String(http.StatusOK, err.Error())
			c.Abort()
		}
		c.String(http.StatusOK, token+"---------------<br>")
		res, err := j.ParseToken(token)
		if err != nil {
			if err == jwtauth.TokenExpired {
				newToken, err := j.RefreshToken(token)
				if err != nil {
					c.String(http.StatusOK, err.Error())
				} else {
					c.String(http.StatusOK, newToken)
				}
			} else {
				c.String(http.StatusOK, err.Error())
			}
		} else {
			c.JSON(http.StatusOK, res)
		}
	})
	authorize := r.Group("/", jwtauth.JWTAuth())
	{
		authorize.GET("user", func(c *gin.Context) {
			claims := c.MustGet("claims").(*jwtauth.CustomClaims)
			fmt.Println(claims.ID)
			c.String(http.StatusOK, claims.Name)
		})
	}

	return r
}
