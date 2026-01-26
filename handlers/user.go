package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"test/models"

	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен. Используйте POST", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса"+ err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Ошибка парсинга JSON"+ err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.DB.Create(&user).Error; err != nil {
		http.Error(w, "Ошибка создания пользователя" +err.Error(), http.StatusBadRequest )
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	 if err :=json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Ошибка кодирования ответа", http.StatusBadRequest)
		return
	}

	 fmt.Printf("Создан новый пользователь: %+v\n", user)
}