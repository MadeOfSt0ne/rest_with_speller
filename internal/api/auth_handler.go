package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"note/internal/auth"
	"note/internal/types"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	authSrv auth.AuthService
}

func NewAuthHandler(authSrv auth.AuthService) *AuthHandler {
	return &AuthHandler{authSrv: authSrv}
}

func (h *AuthHandler) RegisterAuth(r *chi.Mux) {
	r.Post("/login", h.singIn)
}

// Логин в приложении. Принимает в теле запроса login и password. Возвращает jwt токен
func (h *AuthHandler) singIn(w http.ResponseWriter, req *http.Request) {
	logrus.Info("singing in")
	var loginInfo types.LoginInfo
	json.NewDecoder(req.Body).Decode(&loginInfo)
	user, err := h.authSrv.GetUserByLoginAndPassword(loginInfo.Login, loginInfo.Password)
	logrus.Info(user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	token, err := h.authSrv.GenerateToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("oops somethingg went wrong"))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, token)
}
