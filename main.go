package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"test/handlers"
	"test/models"
)

var db *gorm.DB





func init(){
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Ошибка")
	} else {
		fmt.Println("база данных создана")
	}

	db.AutoMigrate(&models.User{}, &models.Event{})
}

func main (){

	eventHandler := &handlers.EventHandler{DB: db}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if r.Method != http.MethodGet {
			http.Error(w, "Ошибка обработки запроса", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "hello go")
		
	})


	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request){

		if r.Method != http.MethodPost {
			http.Error(w, "Ошибка обработки запроса", http.StatusBadRequest)
		}


		fmt.Fprintf(w, "create" )
		body,err := io.ReadAll(r.Body)
		if err != nil{
			http.Error(w, "Ошибка чтения тела запроса",http.StatusBadRequest)
			return
		}

		var user models.User
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Ошибка парсинга JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Ошибка создания пользователя"+ err.Error(), http.StatusInternalServerError )
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(user)

		

		fmt.Println(user)
		


	})


	http.HandleFunc("/event/create", func(w http.ResponseWriter, r *http.Request) {
		eventHandler.CreateEvent(w, r)
	})

	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
		eventHandler.GetEvent(w,r)
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		eventHandler.GetAllEvents(w, r)
	})

	http.ListenAndServe(":8080",nil)
}