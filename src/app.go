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
	http.HandleFunc("/syncDate", handlers.SyncDate)

	url := os.Getenv("APP_URL")
	port := os.Getenv("APP_PORT")

	fmt.Printf("Listening on %s:%s\n", url, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", url, port), nil))
}
