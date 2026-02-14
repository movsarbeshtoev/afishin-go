package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"test/models"

	"gorm.io/gorm"
)


type EventHandler struct {
	DB *gorm.DB
}

type EventInput struct {
	Status string `json:"status"`
}

func pathId (w http.ResponseWriter, r *http.Request) ( eventId uint){
		path := strings.TrimPrefix(r.URL.Path, "/")
	if !strings.HasPrefix(path, "event/") {
		http.Error(w, "Используйте URL вида /event/{id}", http.StatusBadRequest)
		return
	}

	eventID, err := extractIDFromPath(path)
	if err != nil {
		http.Error(w, "Неверный формат ID в URL: "+err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Идет поиск по id:%v\n", eventID)

	

	return eventID

	
}



func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен. Используйте POST", http.StatusBadRequest)
		
		return
	}

	body, err := io.ReadAll(r.Body)
	 if err != nil {
		http.Error(w, "Не удалось прочитать запрос"+err.Error() , http.StatusBadRequest)
		
		return
	 }

	defer r.Body.Close()

	var event models.Event

	if err := json.Unmarshal(body, &event); err != nil {
		http.Error(w, "Ошибка парсинга JSON"+err.Error(), http.StatusBadRequest)
		
		return
	}

	if err := h.DB.Create(&event).Error; err != nil {
		http.Error(w, "Ошибка создания события"+ err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(event); err != nil {
		http.Error(w, "Ошибка кодировки ответа"+err.Error(), http.StatusBadRequest)
		
		return
	}

	fmt.Printf("Событие создано :%+v\n", event )
}

func extractIdFromPath(path string) (uint,error) {
	// убираем начальный и конечный слэш
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")

	// Ожидаем формат: event/{id}

	if len(parts) >= 2 && parts[0] == "event" {
		id, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			return 0, fmt.Errorf("Неверный формат ID: %v" , err)
		}
		
		return uint(id), nil
	}

	return 0, fmt.Errorf("ID не найден в пути")


}

// extractIDFromPath извлекает ID из URL пути вида /event/123
func extractIDFromPath(path string) (uint, error) {
	// Убираем начальный и конечный слэш
	path = strings.Trim(path, "/")
	parts := strings.Split(path, "/")
	
	// Ожидаем формат: event/{id}
	if len(parts) >= 2 && parts[0] == "event" {
		id, err := strconv.ParseUint(parts[1], 10, 32)
		if err != nil {
			return 0, fmt.Errorf("неверный формат ID: %v", err)
		}
		return uint(id), nil
	}
	
	return 0, fmt.Errorf("ID не найден в пути")
}


func (h *EventHandler) GetEvent(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен. Использыйте метод GET", http.StatusMethodNotAllowed)

		return
	}

	var foundEvent models.Event

	eventID := pathId(w, r)

	if err := h.DB.First(&foundEvent, eventID).Error; err != nil {
		http.Error(w, "Событие не  найдено", http.StatusNotFound)

		return
	}

	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&foundEvent); err != nil {
		http.Error(w, "Ошибка кодировки ответа"+ err.Error(), http.StatusInternalServerError)
	}





}

func (h *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {

	
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен. Используйте метод GET", http.StatusMethodNotAllowed)
		
		return
	}

		var events []models.Event
		query := h.DB

		statusParam := r.URL.Query().Get("status")

fmt.Printf("statusParam: %s\n", statusParam)

		switch statusParam {
		case models.EventStatusPending, models.EventStatusCancelled, models.EventStatusCompleted, models.EventStatusPublished :
			query = query.Where("status = ?", statusParam)
		case "":
			
			
		default:
			http.Error(w, "Недопустимый статус", http.StatusBadRequest)	
			return
		}

		categoryIDParam := r.URL.Query().Get("category_id")
		slugParam := r.URL.Query().Get("category")

		if categoryIDParam != "" {
			id,err := strconv.ParseUint(categoryIDParam, 10,32)
			if err != nil{
				http.Error(w, "Недопустимый category_id", http.StatusBadRequest)
				return
			}

			query = query.Where("category_id = ?", id)
		}else if slugParam != ""{
			query =query.Where("category_id IN (SELECT id FROM categories WHERE slug = ?)", slugParam)
		}

		if err := query.Find(&events).Error; err != nil {
			http.Error(w, "Ошибка получения событий из базы данных"+ err.Error(), http.StatusInternalServerError)
			return
		}
	
		// if err := h.DB.Find(&events).Error; err != nil {
		// 	http.Error(w, "Ошибка получения события из базы данных"+err.Error(), http.StatusInternalServerError)

		// 	return
		// }
	

		

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "Ошибка кодировки ответа:"+err.Error(), http.StatusInternalServerError)

		return
	}

	
}


func (h * EventHandler) SetEventStatus(w http.ResponseWriter, r *http.Request ){

	if r.Method != http.MethodPatch {
		http.Error(w, "Этот метод не разрешен, используйте PATCH", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body) 
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	fmt.Printf("body: %s\n", body)

	var input EventInput

	if err := json.Unmarshal(body, &input); err != nil {
		http.Error(w, "Ошибка парсинга json", http.StatusBadRequest)
		return
	}

	eventID := pathId(w, r)

	var event models.Event

	if err := h.DB.First(&event, eventID).Error; err != nil {
		http.Error(w, "Событие не найдено", http.StatusBadRequest)
		return
	}



	switch input.Status {
	case models.EventStatusCancelled,  models.EventStatusCompleted, models.EventStatusPending, models.EventStatusPublished:
		h.DB.Model(&event).Update("status", input.Status)
	default:
		http.Error(w, "Недопустимая роль", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&event); err != nil {
		http.Error(w, "Ошибка кодировки отвеета", http.StatusServiceUnavailable)
	}
}





