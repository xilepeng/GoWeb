package main

import (
	"demo10-2/models"
	"demo10-2/routers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{"UnixToTime": models.UnixToTime})
	r.LoadHTMLGlob("templates/**/*")
	r.Static("/static", "./static")

	routers.DefaultRoutersInit(r) // Gin 路由文件抽离
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)

	err := r.Run()
	if err != nil {
		return
	}
}
