package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type UploadResponse struct {
	Filename     string `json:"filename"`
	OriginalName string `json:"original_name"`
	Size         int64  `json:"size"`
	URL          string `json:"url"`
}


var uploadDir = "./uploads"
const maxUploadSize = 10 << 20 // 10 MB



 func UploadFile (w http.ResponseWriter, r *http.Request){
	r.Body = http.MaxBytesReader(w, r.Body , maxUploadSize)


	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "Фаил слишком большой", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Printf("FormFile error: %v", err)
		http.Error(w, "Ошибка получения файла", http.StatusBadRequest)
		return
	}

	defer file.Close()

	ext := filepath.Ext(handler.Filename)
	allowedExts := map[string]bool{
		".jpg": true, ".jpeg": true, ".png": true,
	}

	if !allowedExts[ext]{
		http.Error(w, "Непотдерживаемый тип файла", http.StatusBadRequest)
		return
	}

	 err = os.MkdirAll(uploadDir, os.ModePerm) 
	 if err != nil{
		http.Error(w, "Ошибка создания дериктории"+ err.Error(), http.StatusBadRequest)

	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)


	dst,err := os.Create(filePath)
	if err !=nil {
		http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	io.Copy(dst, file)


	w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(UploadResponse{
    Filename:     filename,
	OriginalName: handler.Filename,
	Size:         handler.Size,
	URL:          "/uploads/" + filename,
})
}
