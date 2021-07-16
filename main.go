package main

import (
	database "golang-jwt/src/Database"
	"golang-jwt/src/auth"
	"golang-jwt/src/user/infrastructure"
	"golang-jwt/src/user/repository"
	"golang-jwt/src/user/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	db, _ := database.GetConnection()
	repo := repository.NewUserRepo(db)
	userService := service.NewUserService(repo)
	userHandler := infrastructure.NewUserHandler(userService, auth.NewService())
	r.POST("/ping", userHandler.Login)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
