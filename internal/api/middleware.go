package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

var (
	secret = []byte("KodeEducation")
)

type contextKey int

const AuthenticatedUserId contextKey = 0

// Верификация jwt и парсинг id пользователя
func verifyToken(tokenString string) (int, error) {
	logrus.Info("verifying token: ", tokenString)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		logrus.Info("err while parsing token: ", err)
		return -1, err
	}
	if !token.Valid {
		logrus.Info("token is invalid: ", token)
		return -1, fmt.Errorf("invalid token")
	}

	res, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		logrus.Info("failed to typecast to jwt.MapClaims")
		return -1, fmt.Errorf("invalid token")
	}

	userIdRaw := res["id"]
	userId, ok := userIdRaw.(float64)
	if !ok {
		logrus.Info("failed to typecast id to float64")
		return -1, fmt.Errorf("invalid token")
	}
	return int(userId), nil
}

// Валидация токена и добавление id пользователя в контекст запроса
func validateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Missing authorization header")
			return
		}
		tokenString = tokenString[len("Bearer "):]

		userId, err := verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid token")
			return
		}

		ctxWithUserId := context.WithValue(r.Context(), AuthenticatedUserId, userId)
		rWithUser := r.WithContext(ctxWithUserId)
		next.ServeHTTP(w, rWithUser)
	})
}
