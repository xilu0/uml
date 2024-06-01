package main

import (
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// User 结构体
type User struct {
	Name  string `json:"name" binding:"required" validate:"min=3,max=20"`
	Email string `json:"email" binding:"required" validate:"email"`
	Age   int    `json:"age" binding:"required" validate:"min=1,max=100"`
}

func isValidUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	for _, char := range username {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func handleError(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var errMsgs []string
		for _, e := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf("Field: %s, Error: %s", e.Field(), e.Tag()))
		}
		return strings.Join(errMsgs, ", ")
	}
	return err.Error()
}

func main() {
	r := gin.Default()
	validate := validator.New()
	validate.RegisterValidation("username", isValidUsername)

	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully!", "user": user})
	})

	r.Run()
}
