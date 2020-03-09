package main

import (
	"fmt"
	"net/http"
	"rico-ser-gin/routers"
)

func main() {
	routersInit := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 9000),
		Handler:        routersInit,
	}

	server.ListenAndServe()
}