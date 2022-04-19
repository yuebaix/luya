package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/yuebaix/luya/web"
	"html/template"
	"io"
	"net/http"
	"os"
)

func StartAdminServer() {
	engine := initEngine()

	//设置模板
	t, _ := template.ParseFS(web.Template, "template/*.html")
	engine.SetHTMLTemplate(t)
	//设置静态资源
	engine.StaticFS("/ui", http.FS(web.Static))

	//配置路由
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{"title": "Luya Template"})
	})

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "farewell")
	})

	engine.Run()
}

func initEngine() *gin.Engine {
	//设置日志流
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//设置日志格式
	gin.DisableConsoleColor()

	//路由格式
	/*gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+"%-6s %-25s --> %s (%d handlers)\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}*/

	router := gin.Default()

	//请求格式
	/*router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())*/

	return router
}
