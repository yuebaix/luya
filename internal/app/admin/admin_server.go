package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/yuebaix/luya/web"
	"html/template"
	"net/http"
)

func StartAdminServer() {
	t, _ := template.ParseFS(web.Template, "template/*.html")

	router := gin.Default()
	router.SetHTMLTemplate(t)

	router.StaticFS("/ui", http.FS(web.Static))

	router.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"title": "Luya Template"})
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "farewell")
	})

	router.Run()
}
