package auth

import (
	"CS230832/BeastVehicles/config"
	"CS230832/BeastVehicles/types"
	"CS230832/BeastVehicles/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJWT(secret string, email string) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(expiration).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.AdminStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := utils.GetTokenFromRequest(r)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)

		admin, err := store.GetAdmin(email)
		if err != nil {
			permissionDenied(w)
			return
		}

		loginTokens, err := store.GetTokens(email)
		if err != nil {
			permissionDenied(w)
			return
		}

		hasLoggedIn := false
		userAgent := r.Header.Get("User-Agent")
		
		for _, loginToken := range loginTokens {
			if loginToken == userAgent {
				hasLoggedIn = true
				break
			}
		}

		if !hasLoggedIn {
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = utils.NewContext(ctx, admin)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}
