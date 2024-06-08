package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomRequest struct {
	StringField  string  `form:"string_field"`
	IntField     int     `form:"int_field"`
	BoolField    bool    `form:"bool_field"`
	RuneField    rune    `form:"rune_field"`
	PointerField *string `form:"pointer_field"`
}

func main() {
	r := gin.Default()

	r.POST("/custom-request", func(c *gin.Context) {
		var req CustomRequest
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, req)
	})

	r.Run(":8080")
}
