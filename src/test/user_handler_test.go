package test

import (
	"fmt"
	database "golang-jwt/src/Database"
	"golang-jwt/src/auth"
	"golang-jwt/src/user/infrastructure"
	"golang-jwt/src/user/repository"
	"golang-jwt/src/user/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Test_login_succeess(t *testing.T) {

	r := gin.Default()
	db, _ := database.GetConnection()
	repo := repository.NewUserRepo(db)
	userService := service.NewUserService(repo)
	userHandler := infrastructure.NewUserHandler(userService, auth.NewService())
	r.POST("/ping", userHandler.Login)

	requestBody := strings.NewReader("username=ariadi&password=12345")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/ping", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
