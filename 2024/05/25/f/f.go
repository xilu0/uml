package main

import (
	"fmt"
	"net/http"
)

// 控制器层
func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	greeting := getGreeting(name)
	fmt.Fprintf(w, greeting)
}

// 服务层
func getGreeting(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
