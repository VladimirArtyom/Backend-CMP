package main

import (
	"compro-exh/internal/db"
	"fmt"
	"log"
	"net/http"
)

func main() {

	db.Init()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", 9999),
		Handler: routes(),
	}

	err := srv.ListenAndServe()

	fmt.Printf("Listen and serve")

	if err != nil {
		log.Panic(err)
	}
}
