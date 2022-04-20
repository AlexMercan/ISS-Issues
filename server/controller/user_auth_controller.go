package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"server/model"
	"server/repository"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type UserAuthController struct {
	repo   repository.IUserAuthRepository
	secret []byte
}

func CreateUserAuthController(repo repository.IUserAuthRepository, secret []byte) *UserAuthController {
	return &UserAuthController{repo, secret}
}

func (authController *UserAuthController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user := authController.repo.FindOneByUsername(credentials.Username)
	err = bcrypt.CompareHashAndPassword([]byte(user.Credentials.Password), []byte(credentials.Password))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(authController.secret)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expirationTime,
	})
	w.WriteHeader(http.StatusOK)
}

func (authController *UserAuthController) Register(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	credentials.Password = string(hashedPass)
	_, err = authController.repo.Save(model.User{ID: 0, Credentials: credentials})
	if err != nil {
		log.Print(err)
		repoError := ErrorResponse{Message: "Username already exists"}
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(repoError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
