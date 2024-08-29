package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"note/internal/types"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type AuthSrv struct {
}

func NewAuthService() AuthSrv {
	return AuthSrv{}
}

var (
	secret = []byte("KodeEducation")
)

func (s *AuthSrv) GetUserByLoginAndPassword(login, password string) (types.User, error) {
	logrus.Info("looking for user")
	logrus.Infof("Getting user by unique login: %s", login)

	users := map[string]types.User{
		"artur": {
			Id:       4,
			Login:    "archi",
			Password: "d74ff0ee8da3b9806b18c877dbf29bbde50b5bd8e4dad7a3a725000feb82e8f1"}, // pass
		"ruslan": {
			Id:       7,
			Login:    "rusya",
			Password: "98c1eb4ee93476743763878fcb96a25fbc9a175074d64004779ecb5242f645e6"}, // word
	}

	user, ok := users[login]
	if !ok {
		return types.User{}, fmt.Errorf("user not found")
	}

	if !checkPasswordHash(password, user.Password) {
		return types.User{}, fmt.Errorf("login or password was incorrect")
	}
	return user, nil
}

func checkPasswordHash(password, passwordFromDB string) bool {
	hash := sha256.Sum256([]byte(password))
	hashString := hex.EncodeToString(hash[:])

	return hashString == passwordFromDB
}

func (s *AuthSrv) GenerateToken(user types.User) (string, error) {
	logrus.Info("Generating jwt")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"login": user.Login,
			"id":    user.Id,
		})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
