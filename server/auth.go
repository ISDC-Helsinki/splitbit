package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"isdc.fi/splitbit/server/api"
	"isdc.fi/splitbit/server/data"
)

type Security struct{}
type CustomClaims struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var secretKey = []byte("secret key")

func GenerateJWT(userID string) (string, error) {
	claims := CustomClaims{
		UserID: 1,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1 day expiration
			Issuer:    "your-issuer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
func VerifyJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm (optional)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (h *Security) HandleBearerAuthCookie(ctx context.Context, operationName string, t api.BearerAuthCookie) (context.Context, error) {

	claims, err := VerifyJWT(t.GetAPIKey())
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, "user_id", int64(claims.UserID))
	return ctx, nil
}
func (h *Security) HandleBearerAuthHeader(ctx context.Context, operationName string, t api.BearerAuthHeader) (context.Context, error) {

	claims, err := VerifyJWT(t.GetToken())
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, "user_id", int64(claims.UserID))
	return ctx, nil
}

func (h *Handler) LoginPost(ctx context.Context, req *api.UserCredentials) (api.LoginPostRes, error) {
	user, err := qs.GetUserByUsernameAndPassword(ctx, data.GetUserByUsernameAndPasswordParams{Username: req.Username, Password: req.Password})

	if err != nil {
		return nil, errors.New("Incorrect credentials")
	}
	jwt, err := GenerateJWT(strconv.Itoa(int(user.ID)))
	if err != nil {
		resp := api.LoginPostUnauthorizedApplicationJSON("Wrong credentials.")
		return &resp, errors.New("Incorrect credentials")
	}
	c := http.Cookie{
		Name:     "jwt",
		Value:    url.QueryEscape(jwt),
		MaxAge:   3600,
		Path:     "/",
		Domain:   "",
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
		HttpOnly: true,
	}
	return &api.LoginSuccessHeaders{
		SetCookie: c.String(), Response: api.LoginSuccess{Token: jwt},
	}, nil
}

func (h *Handler) RegisterPost(ctx context.Context, req *api.RegisterPostReq) (api.RegisterPostRes, error) {
	_, err := qs.AddUser(ctx, data.AddUserParams{Username: req.Username, Displayname: req.Username, Password: req.Password})

	if err != nil {
		println(err)
		return nil, errors.New("Db insertion error")
	}
	return &api.RegisterPostOK{}, nil
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		log.Printf("Method: %s, Path: %s", r.Method, r.URL.String())
		// Pass to the next handler if it's not a preflight request
		next.ServeHTTP(w, r)
	})
}
