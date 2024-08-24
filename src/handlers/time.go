package handlers

import (
	"net/http"
	"os"
	"time"
)

func SyncDate(w http.ResponseWriter, _ *http.Request) {
	loc, _ := time.LoadLocation(os.Getenv("APP_TIMEZONE"))
	time.Local = loc

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(time.Now().Format("2006-01-02")))
}
