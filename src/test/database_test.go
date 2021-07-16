package test

import (
	"fmt"
	database "golang-jwt/src/Database"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := database.GetConnection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
