package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			logger.Printf("ParseMultipartForm error: %v", err)
			http.Error(w, "File too large", http.StatusBadRequest)
			return
		}

		file, header, err := r.FormFile("myFile")
		if err != nil {
			logger.Printf("Error retrieving file: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		defer file.Close()

		logger.Printf("Received file: %s (Size: %d bytes)", header.Filename, header.Size) // Добавлено логирование

		if header.Size == 0 {
			logger.Println("Empty file uploaded")
			http.Error(w, "Empty file", http.StatusBadRequest)
			return
		}

		data, err := io.ReadAll(file)
		if err != nil {
			logger.Printf("Error reading file: %v", err)
			http.Error(w, "Internal error", http.StatusInternalServerError)
			return
		}

		result, err := service.ConvStr(string(data))
		if err != nil {
			logger.Printf("Conversion error: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ext := filepath.Ext(header.Filename)
		timestamp := time.Now().UTC().Format("2006-01-02T15:04:05Z")
		filename := fmt.Sprintf("%s%s",
			strings.ReplaceAll(timestamp, ":", "-"),
			ext)

		if err := os.WriteFile(filename, []byte(result), 0644); err != nil {
			logger.Printf("File save error: %v", err)
			http.Error(w, "Result save failed", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}
