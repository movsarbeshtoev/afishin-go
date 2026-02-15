package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"test/models"

	"gorm.io/gorm"
)

type CategoryHandler struct {
	DB *gorm.DB
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Этот метод не разрешен, используйте POST", http.StatusMethodNotAllowed)
		return
	}

	body,err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var cat models.Category

	if err := json.Unmarshal(body, &cat); err != nil {
		http.Error(w, "Ошибка парсинга json", http.StatusBadRequest)
		return
	}

	if cat.Name == ""{
		http.Error(w, "Категория пуста", http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&cat).Error; err != nil {
		http.Error(w, "Ошибка создания категории" + err.Error(), http.StatusInternalServerError)
        return
	}

	w.Header().Set("Content-Type", "applicatio/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(cat)

}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
       http.Error(w, "Этот метод не разрешен, используйте POST", http.StatusMethodNotAllowed)
		return
    }


    var categories []models.Category

    if err := h.DB.Find(&categories).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(categories)
}