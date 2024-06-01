package main

import (
	"fmt"
)

func main() {
	dsn := "user:password@/dbname"
	server := InitializeServer(dsn)
	fmt.Println("Server initialized with database:", server.DB.DSN)
}
