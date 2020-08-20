package main

import (
	"fmt"
	"net/http"

	"github.com/fedorkolmykow/messesnger/pkg/dbclient"
	"github.com/fedorkolmykow/messesnger/pkg/service"
	"github.com/fedorkolmykow/messesnger/pkg/service/httpserver"
)

func main() {
	dbCon := dbclient.NewDb()
	svc := service.NewService(dbCon)
	server := httpserver.NewServer(svc)

	fmt.Println("starting server at :9000")
	err := http.ListenAndServe(":9000", server)
	if err != nil {
		fmt.Println(err)
	}
}