package main

import (
	"context"
	"fmt"
	"net/http"

	"avito_message/message_service/pkg/dbclient"
	"avito_message/message_service/pkg/service"
	"avito_message/message_service/pkg/service/httpserver"
)

func main() {
	dbCon := dbclient.NewDb(context.Background())
	service := service.NewService(dbCon)
	server := httpserver.NewServer(service)

	fmt.Println("starting server at :9000")
	err := http.ListenAndServe(":9000", server)
	if err != nil {
		fmt.Println(err)
	}
}