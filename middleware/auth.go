package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"test/models"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

// contextKey — свой тип ключа контекста, чтобы не пересекаться с другими пакетами.
type contextKey string

const UserIDKey contextKey = "user_id"
const RoleKey   contextKey = "role"

// Claims должен совпадать с тем, что мы кладём в токен в handlers/auth.go.
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// AuthMiddleware читает JWT_SECRET из переменной окружения.
func AuthMiddleware() func(http.Handler) http.Handler {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Ожидаем заголовок: Authorization: Bearer <token>
			auth := r.Header.Get("Authorization")
			if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
				http.Error(w, "требуется авторизация: заголовок Authorization: Bearer <token>", http.StatusUnauthorized)
				return
			}
			tokenString := strings.TrimPrefix(auth, "Bearer ")

			var claims Claims
			token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
				return jwtSecret, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "недействительный или истёкший токен", http.StatusUnauthorized)
				return
			}

			// Кладём user_id в контекст — в handler можно взять: middleware.UserIDFromRequest(r)
			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// UserIDFromRequest извлекает user_id из контекста (после AuthMiddleware).
func UserIDFromRequest(r *http.Request) (uint, bool) {
	id, ok := r.Context().Value(UserIDKey).(uint)
	return id, ok
}

// RoleFromRequest извлекает роль из контекста (если задана в AuthMiddleware).
func RoleFromRequest(r *http.Request) (string, bool) {
	role, ok := r.Context().Value(RoleKey).(string)
	return role, ok
}

func RequireRole(db *gorm.DB, allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userID, ok := UserIDFromRequest(r)
			if !ok {
				http.Error(w, "требуется авторизация", http.StatusUnauthorized)
				return
			}

			var user models.User
			if err := db.First(&user, userID).Error; err != nil {
				http.Error(w, "пользователь не найден", http.StatusUnauthorized)
				return
			}

			for _, allowed := range allowedRoles {
				if user.Role == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "доступ запрещён", http.StatusForbidden)
		})
	}
}