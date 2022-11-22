package database

import "fmt"

var connection string

func init() {
	fmt.Println("init executed")
	connection = "PostgreSQL"
}

func GetDatabase() string {
	return connection
}
