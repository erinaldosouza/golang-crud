package routes

import (
	"fmt"
	"golang-crud/app/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:id", func(ctx *gin.Context) { fmt.Println("called the get handler") }, controller.FindUserById)
	r.POST("/user", func(ctx *gin.Context) { fmt.Println("called the post handler") }, controller.CreateUserById)
	r.PUT("/user/:id", func(ctx *gin.Context) { fmt.Println("called the put handler") }, controller.UpdateUserById)
	r.DELETE("/user/:id", func(ctx *gin.Context) { fmt.Println("called the delete handler") }, controller.DeleteUserById)
}
