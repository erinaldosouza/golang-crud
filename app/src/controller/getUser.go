package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	fmt.Println("called the get controller")
}
