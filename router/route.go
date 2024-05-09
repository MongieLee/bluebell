package router

import (
	"bluebell/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	}
	r := gin.New()
	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404",
		})
	})
	return r
}
