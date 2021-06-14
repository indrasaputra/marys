package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/indrasaputra/marys"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is not set")
	}

	http.HandleFunc("/notifications", marys.ReceiveNotification)
	_ = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
