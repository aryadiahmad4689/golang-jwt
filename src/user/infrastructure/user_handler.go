package infrastructure

import (
	"fmt"
	"golang-jwt/src/auth"
	"golang-jwt/src/user/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserServiceInterface
	auth        auth.Service
}

type Json struct {
	Token string `json:"token"`
}

func NewUserHandler(userService service.UserServiceInterface, auth auth.Service) *UserHandler {
	return &UserHandler{userService: userService, auth: auth}
}

func (UserHandler *UserHandler) Login(c *gin.Context) {
	username := c.PostForm("username")
	pass := c.PostForm("password")

	result, err := UserHandler.userService.Login(username, pass)
	if err != nil {
		// json.Unmarshal([]byte())
		fmt.Println(err)
		return
	}

	token, _ := UserHandler.auth.GenerateToken(result.ID)
	json := Json{Token: token}

	fmt.Println(token)
	c.JSON(200, json)

}
