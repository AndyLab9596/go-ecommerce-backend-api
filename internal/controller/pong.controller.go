package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	name := c.Param("name")
	uid := c.Query("uid")
	age := c.DefaultQuery("age", "30")

	ageInt, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("convert age error::", err)
		ageInt = 0
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pong " + name,
		"uid":     uid,
		"age":     ageInt,
		"users":   []string{"cr7", "m10"},
	})
}
