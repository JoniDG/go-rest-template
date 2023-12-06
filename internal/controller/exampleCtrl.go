package controller

import (
	"github.com/JoniDG/go-rest-template/internal/service"
	"github.com/gin-gonic/gin"
)

type ExampleController interface {
	ExampleHandler(ctx *gin.Context)
}
type exampleController struct {
	svc service.ExampleService
}

func NewExampleController(svc service.ExampleService) ExampleController {
	return &exampleController{svc: svc}
}

func (c *exampleController) ExampleHandler(ctx *gin.Context) {

}
