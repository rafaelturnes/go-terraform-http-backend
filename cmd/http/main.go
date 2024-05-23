package main

import "go-terraform-http-backend/internal/adapter/handler/http"

func main() {
	server := http.NewServer()
	server.Routes()
	server.Run()
}
