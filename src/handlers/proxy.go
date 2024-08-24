package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/salahfarzin/api/src/types"
	"github.com/salahfarzin/api/src/utils"
	"github.com/salahfarzin/utils-go/logger"
	"go.uber.org/zap/zapcore"
)

const (
	logResponseKey = "Response"
	contentType    = "application/json"
	envTmpPath     = "TMP_PATH"
)

func Proxy(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)

	logger.Info(r.URL.Path, zapcore.Field{
		Key:    logRequestKey,
		Type:   zapcore.FieldType(zapcore.StringType),
		String: r.Method,
	})

	response := makeRequest(r)

	if err := prepareStorageDirectories(); err != nil {
		panic(err)
	}

	// map download links to handle by proxy
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) >= 3 {
		methodName := segments[2]

		switch methodName {
		case "downloadPdf", "getUpdates":
			var responseItems = response.Items.(map[string]any)
			url := responseItems["link"].(string)
			responseItems["link"] = mapUrl(url)

			download(url)
			response.Items = responseItems
		case "questions":
			downloadQuestionImages(&response)
		}
	}

	status := "Error"
	if response.Success {
		status = "Success"
	}

	logger.Info(r.URL.Path, zapcore.Field{
		Key:       logResponseKey,
		Type:      zapcore.FieldType(zapcore.StringType),
		String:    status,
		Interface: response,
	})

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func makeRequest(r *http.Request) types.RemoteResponse {
	var response types.RemoteResponse

	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("read request body", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		var formData = make(map[string]any)
		for fieldName := range r.Form {
			formData[fieldName] = r.FormValue(fieldName)
		}

		body, err = json.Marshal(formData)
		if err != nil {
			logger.Error("json marshal", err)
		}
	}

	client := &http.Client{
		Timeout: 300 * time.Second,
	}

	req, err := http.NewRequest(r.Method, os.Getenv("REMOTE_URL")+r.URL.Path, bytes.NewBuffer(body))
	if err != nil {
		logger.Error("remote server", err)
	}

	if os.Getenv("APP_ENV") == "dev" {
		req.Header.Add("Cookie", "XDEBUG_SESSION=PHPSTORM")
	}

	req.Header.Add("Accept", contentType)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", r.Header.Get("Authorization"))

	res, err := client.Do(req)
	if err != nil {
		logger.Error("trying to access remote server", err)
	}
	defer res.Body.Close()

	response.StatusCode = res.StatusCode
	json.NewDecoder(res.Body).Decode(&response)

	return response
}

func download(url string) {
	segments := strings.Split(url, "/")
	fileName := fmt.Sprintf("%s-%s", segments[len(segments)-2], segments[len(segments)-1])

	err := utils.Download(url, fmt.Sprintf("%s/%s", os.Getenv(envTmpPath), fileName))
	if err != nil {
		logger.Error(err.Error(), err)
	}
}

func downloadQuestionImages(response *types.RemoteResponse) error {
	var questions []types.Question

	byteItems, err := json.Marshal(response.Items)
	if err != nil {
		return err
	}

	json.Unmarshal(byteItems, &questions)
	for qi, q := range questions {
		if len(q.ImgLink) > 0 {
			download(q.ImgLink)

			q.ImgLink = mapUrl(q.ImgLink)
		}

		for ai, a := range q.Answers {
			if len(a.ImgLink) == 0 {
				continue
			}

			download(a.ImgLink)
			q.Answers[ai].ImgLink = mapUrl(a.ImgLink)
		}

		questions[qi] = q
	}

	response.Items = questions

	return nil
}

func mapUrl(url string) string {
	segments := strings.Split(url, "/")

	var newUrl string = fmt.Sprintf("%s:%s", os.Getenv("CLIENT_URL"), os.Getenv("CLIENT_PORT"))
	for i, v := range segments {
		if i <= 2 {
			continue
		}
		newUrl += "/" + v
	}

	return newUrl
}

func prepareStorageDirectories() error {
	path := os.Getenv(envTmpPath)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	return nil
}
