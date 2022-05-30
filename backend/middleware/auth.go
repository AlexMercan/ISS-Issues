package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/controller"

	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("token")
		log.Printf("Cookies: %v", r.Cookies())
		if err != nil {
			log.Printf("Error: %s", err)
			noCookieToken := controller.ErrorResponse{Message: err.Error()}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(noCookieToken)
			return
		}
		token, err := jwt.ParseWithClaims(tokenCookie.Value, &controller.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				e := fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
				log.Printf("Unexpected signing method: %x", token.Header["alg"])
				return nil, e
			}
			return []byte("verysecretkey"), nil
		})
		if claims, ok := token.Claims.(*controller.Claims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), "props", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			log.Println(err)
			unauthorizedError := controller.ErrorResponse{Message: "Unauthorized"}
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(unauthorizedError)
		}
	})
}
