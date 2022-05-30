package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"server/ent"
	"server/ent/user"
	"server/repository"
	"time"
	"unicode/utf8"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string    `json:"username"`
	Role     user.Role `json:"role"`
	Id       int       `json:"id"`
	jwt.RegisteredClaims
}

type UserAuthController struct {
	repo   *repository.UserAuthRepository
	secret []byte
}

func CreateUserAuthController(repo *repository.UserAuthRepository, secret []byte) *UserAuthController {
	return &UserAuthController{repo, secret}
}

func (authController *UserAuthController) Login(w http.ResponseWriter, r *http.Request) {
	var userCredentials ent.User
	err := json.NewDecoder(r.Body).Decode(&userCredentials)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := authController.repo.FindOneByUsername(r.Context(), userCredentials.Username)
	if err != nil {
		log.Print(err)
		repoError := ErrorResponse{Type: "validation", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(repoError)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password))
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(5 * time.Hour)
	claims := &Claims{
		Username: userCredentials.Username,
		Role:     userCredentials.Role,
		Id:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: expirationTime},
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
		Name:     "token",
		Value:    tokenStr,
		Expires:  expirationTime,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	user.Password = ""
	_ = json.NewEncoder(w).Encode(user)
}

func (authController *UserAuthController) Register(w http.ResponseWriter, r *http.Request) {
	var user ent.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if utf8.RuneCountInString(user.Password) < 8 {
		log.Print(err)
		repoError := ErrorResponse{Type: "validation", Message: "Password is too short"}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(repoError)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	hashedPassStr := string(hashedPass)
	_, err = authController.repo.Save(r.Context(), user.Username, hashedPassStr, user.Role)
	if err != nil {
		log.Print(err)
		repoError := ErrorResponse{Type: "validation", Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(repoError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
