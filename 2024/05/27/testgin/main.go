// main.go
package main

import (
	"strconv"

	"github.com/xilu0/uml/2024/05/27/testgin/container"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	c := container.NewContainer()

	r.GET("/user/:id", func(ctx *gin.Context) {
		GetUserHandler(ctx, c)
	})

	r.Run()
}

func GetUserHandler(ctx *gin.Context, userService *container.Container) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := userService.UserService.GetUser(id)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Unable to retrieve user"})
		return
	}

	ctx.JSON(200, user)
}
