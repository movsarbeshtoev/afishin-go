package main

import (
	"fmt"
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

	db.AutoMigrate(&models.User{}, &models.Event{}, &models.Category{})
}

func main (){

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	authHandler := &handlers.AuthHandler{DB: db, JWTSecret: jwtSecret}
	authMiddleware := middleware.AuthMiddleware()

	eventHandler := &handlers.EventHandler{DB: db}

	categoryHanler := &handlers.CategoryHandler{DB: db}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if r.Method != http.MethodGet {
			http.Error(w, "Ошибка обработки запроса", http.StatusBadRequest)
		}

		fmt.Fprintf(w, "hello go")
		
	})



	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)






	

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

	http.HandleFunc("/category", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHanler.GetCategories(w, r)
		case http.MethodPost:
			h := authMiddleware(middleware.RequireRole(db,models.RoleAdmin)(http.HandlerFunc(categoryHanler.CreateCategory)))
			h.ServeHTTP(w, r)
		
		default: 
			http.Error(w, "Метод не разрешен, используйте GET или POST", http.StatusMethodNotAllowed)
		}
		
	})



	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer((http.Dir("uploads")))))

	http.ListenAndServe(":8080",middleware.CORS(http.DefaultServeMux))
}