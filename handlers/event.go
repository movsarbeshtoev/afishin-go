package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"test/models"

	"gorm.io/gorm"
)


type EventHandler struct {
	DB *gorm.DB
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r * http.Request){

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


func (h *EventHandler) GetEvent(w http.ResponseWriter, r * http.Request){

	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен. Использыйте метод GET", http.StatusMethodNotAllowed)

		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Не указан ID", http.StatusBadRequest)
		return
	}

	var eventID uint
	if _, err := fmt.Sscanf(idStr, "%d", &eventID); err != nil || eventID == 0 {
		http.Error(w, "Неверный ID", http.StatusBadRequest)
		return
	}

	//  body,err := io.ReadAll(r.Body); 
	//  if err != nil {
	// 	http.Error(w, "Ошипка чтения запроса"+err.Error(), http.StatusBadRequest)

	// 	return
	// }

	// r.Body.Close()

	// var event models.Event
	// if err := json.Unmarshal(body,&event); err != nil {
	// 	http.Error(w, "Ошипка парсинга json"+ err.Error(), http.StatusBadRequest)

			
	// 	return
	// }

	fmt.Printf("Идет поиск по id:%v\n", eventID)

	// if event.ID == 0 {
	// 	http.Error(w, "Не указан ID", http.StatusBadRequest)

	// 	return
	// }

	var foundEvent models.Event

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

func (h *EventHandler) GetAllEvents(w http.ResponseWriter, r * http.Request) {

	
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не разрешен. Используйте метод GET", http.StatusMethodNotAllowed)
		
		return
	}

		var events []models.Event
	
		if err := h.DB.Find(&events).Error; err != nil {
			http.Error(w, "Ошибка получения события из базы данных"+err.Error(), http.StatusInternalServerError)

			return
		}
	

		

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(events); err != nil {
		http.Error(w, "Ошибка кодировки ответа:"+err.Error(), http.StatusInternalServerError)

		return
	}

	
}






