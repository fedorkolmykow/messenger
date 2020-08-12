package main

import (
	"fmt"
	"net/http"

	"github.com/fedorkolmykow/messages/pkg/dbclient"
	"github.com/fedorkolmykow/messages/pkg/service"
	"github.com/fedorkolmykow/messages/pkg/service/httpserver"
)

func main() {
	dbCon := dbclient.NewDb()
	service := service.NewService(dbCon)
	server := httpserver.NewServer(service)

	fmt.Println("starting server at :9000")
	err := http.ListenAndServe(":9000", server)
	if err != nil {
		fmt.Println(err)
	}
}