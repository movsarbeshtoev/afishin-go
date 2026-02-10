package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"test/handlers"
	"test/middleware"
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

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	authHandler := &handlers.AuthHandler{DB: db, JWTSecret: jwtSecret}
	authMiddleware := middleware.AuthMiddleware()

	eventHandler := &handlers.EventHandler{DB: db}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if r.Method != http.MethodGet {
			http.Error(w, "Ошибка обработки запроса", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "hello go")
		
	})



	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)





	http.Handle("/create", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

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
		


	})))

	

	http.Handle("/event", authMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		eventHandler.CreateEvent(w, r)
	})))

	// 	http.HandleFunc("/event", func(w http.ResponseWriter, r *http.Request) {
	// eventHandler.CreateEvent(w, r)
	// })


	http.HandleFunc("/event/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			eventHandler.GetEvent(w, r)
		case http.MethodPatch:
			h := authMiddleware(middleware.RequireRole(db, models.RoleAdmin, models.RoleModerator)(http.HandlerFunc(eventHandler.SetEventStatus)))
			h.ServeHTTP(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		eventHandler.GetAllEvents(w, r)
	})

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		handlers.UploadFile(w,r)
	})


	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer((http.Dir("uploads")))))

	http.ListenAndServe(":8080",middleware.CORS(http.DefaultServeMux))
}