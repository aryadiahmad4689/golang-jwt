package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_login_succeess(t *testing.T) {

	// r := gin.Default()
	// db, _ := database.GetConnection()
	// repo := repository.NewUserRepo(db)
	// userService := service.NewUserService(repo)
	// userHandler := infrastructure.NewUserHandler(userService, auth.NewService())

	requestBody := strings.NewReader("username=ariadi&password=12345")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/ping", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
