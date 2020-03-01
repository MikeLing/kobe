package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"kobe/pkg/api/inventory"
	"kobe/pkg/api/playbook"
	"kobe/pkg/api/task"
	"kobe/pkg/api/worker"
	"kobe/pkg/db"
	"kobe/pkg/middlewares"
)

var App *gin.Engine

func init() {
	App = gin.Default()
}

func Run() error {
	db.Connect()
	App.Use(middlewares.Connect)
	App.Use(middlewares.WorkerManager)
	v1 := App.Group("/api/v1")
	{
		p := v1.Group("/playbooks")
		{
			p.GET("/", playbook.List)
		}
		i := v1.Group("/inventory")
		{
			i.GET("/", inventory.List)
			i.POST("/", inventory.Create)
			i.PUT("/:name", inventory.Update)
			i.GET("/:name", inventory.Get)
			i.DELETE("/:name", inventory.Delete)
		}
		t := v1.Group("tasks")
		{
			t.GET("/", task.List)
			t.GET("/:uid", task.Get)
			t.POST("/", task.Create)
		}
		w := v1.Group("/workers")
		{
			w.GET("/",worker.List)
		}
	}
	bind := viper.GetString("server.bind")
	return App.Run(bind)
}