package router

import (
	"net/http"

	"github.com/JoniDG/go-rest-template/internal/controller"
	"github.com/JoniDG/go-rest-template/internal/repository"
	"github.com/JoniDG/go-rest-template/internal/service"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	mapRoutes(r)
	return r
}

func mapRoutes(r *gin.Engine) {

	ejRepo := repository.NewExampleRepository()
	ejSvc := service.NewExampleService(ejRepo)
	ejCtrl := controller.NewExampleController(ejSvc)

	r.GET("/exampleget", ejCtrl.ExampleHandler)
	r.GET("/ping", ping)
}

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}
