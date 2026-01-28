# Работа с файлами в REST API - Полное руководство

## Содержание

1. [Загрузка файлов (Upload)](#загрузка-файлов-upload)
2. [Скачивание файлов (Download)](#скачивание-файлов-download)
3. [Примеры реализации на Go](#примеры-реализации-на-go)
4. [Примеры клиентского кода](#примеры-клиентского-кода)
5. [Best Practices](#best-practices)

---

## Загрузка файлов (Upload)

### 1. Multipart/form-data (рекомендуется)

**Сервер (Go):**
package handlers

import (
"fmt"
"io"
"net/http"
"os"
"path/filepath"
)

const uploadDir = "./uploads"
const maxUploadSize = 10 << 20 // 10 MB

func UploadFile(w http.ResponseWriter, r \*http.Request) {
r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)

    if err := r.ParseMultipartForm(maxUploadSize); err != nil {
        http.Error(w, "Файл слишком большой", http.StatusBadRequest)
        return
    }

    file, handler, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Ошибка получения файла", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Проверка расширения
    ext := filepath.Ext(handler.Filename)
    allowedExts := map[string]bool{
        ".jpg": true, ".jpeg": true, ".png": true, ".gif": true,
        ".pdf": true, ".doc": true, ".docx": true,
    }
    if !allowedExts[ext] {
        http.Error(w, "Неподдерживаемый тип файла", http.StatusBadRequest)
        return
    }

    // Создаем директорию
    os.MkdirAll(uploadDir, os.ModePerm)

    // Генерируем уникальное имя
    filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
    filepath := filepath.Join(uploadDir, filename)

    // Сохраняем файл
    dst, err := os.Create(filepath)
    if err != nil {
        http.Error(w, "Ошибка создания файла", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    io.Copy(dst, file)

    // Возвращаем результат
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "filename": filename,
        "original_name": handler.Filename,
        "size": handler.Size,
        "url": "/uploads/" + filename,
    })

}
