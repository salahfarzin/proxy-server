package src

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/salahfarzin/api/src/handlers"
)

func Start() {
	http.HandleFunc("/", handlers.Proxy)
	http.HandleFunc("/download/", handlers.Download)
	http.HandleFunc("/ping", handlers.Ping)

	port := os.Getenv("APP_PORT")

	fmt.Printf("Listening on 127.0.0.1:%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
