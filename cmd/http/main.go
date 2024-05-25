package main

import (
	"go-terraform-http-backend/internal/adapter/handler/http"
	inmemory "go-terraform-http-backend/internal/adapter/storage/in_memory"
	"go-terraform-http-backend/internal/core/service"
)

func main() {

	stateStorage := inmemory.NewStateStorage()

	stateSVC := service.NewStateService(stateStorage)

	server := http.NewServer(stateSVC)

	server.Routes()

	server.Run()
}
