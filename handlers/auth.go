package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"test/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type RegisterInput struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}

type AuthHandler struct {
	DB *gorm.DB
	JWTSecret []byte
}

type LoginInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User models.User `json:"user"`
}

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// generateJWT создаёт подписанный токен с user_id и сроком жизни (например 24 часа).
func (h *AuthHandler) generateJWT(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.JWTSecret)
}



func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метож не разрешен. Используйте метод POST", http.StatusMethodNotAllowed)
		return
	}

	body,err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения запроса", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var input RegisterInput

	if err := json.Unmarshal(body, &input); err != nil {
		http.Error(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	if input.Email == "" || input.Password == "" || input.Phone == "" || input.Name == "" {
		http.Error(w, "email, password, name и phone обязательны", http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password),bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Ошибка хэширования пароля", http.StatusInternalServerError)
		return
	}


	user := models.User{
		Name: input.Name,
		Email: input.Email,
		Phone: input.Phone,
		PasswordHash: string(hash),
	}

	if err := h.DB.Create(&user).Error; err !=nil {
		http.Error(w, "Ошибка создания пользователя (возможно email или phone уже заняты)", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "app;ication/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Ошибка кодировки ответа"+ err.Error(), http.StatusInternalServerError)
	}

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен. Используйте POST", http.StatusMethodNotAllowed)
		return
	}
		body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Ошибка чтения тела запроса", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var input LoginInput

	if err := json.Unmarshal(body, &input); err != nil{
		http.Error(w, "Ошибка парсинга JSON", http.StatusBadRequest)
		return
	}

	if input.Email == "" || input.Password == "" {
		http.Error(w, "email и password обязательны", http.StatusBadRequest)
		return
	}

	var user models.User

	if err := h.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "неверный email или пароль", http.StatusUnauthorized)
		return
		}

	http.Error(w, "Ошибка базы данных", http.StatusInternalServerError)
	return
		
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil{
		http.Error(w, "неверный email или пароль", http.StatusUnauthorized)
		return
	}

	tokenString, err := h.generateJWT(user.ID)
	if err != nil{
		http.Error(w, "Ошибка создания токена", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: tokenString , User: user})


}