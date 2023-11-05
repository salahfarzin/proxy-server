package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/salahfarzin/utils-go/logger"
	"go.uber.org/zap/zapcore"
)

const logRequestKey = "Request"

func Download(w http.ResponseWriter, r *http.Request) {
	logger.Info(r.RequestURI, zapcore.Field{
		Key:    logRequestKey,
		Type:   zapcore.FieldType(zapcore.StringType),
		String: r.Method,
	})

	segments := strings.Split(r.URL.Path, "/")
	fileName := fmt.Sprintf("%s-%s", segments[len(segments)-2], segments[len(segments)-1])

	http.ServeFile(w, r, fmt.Sprintf("%s/%s", os.Getenv("TMP_PATH"), fileName))
}
