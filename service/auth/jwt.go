package auth

import (
	"CS230832/BeastVehicles/config"
	"CS230832/BeastVehicles/types"
	"CS230832/BeastVehicles/utils"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func NewJWT(secret string, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username})

	return token.SignedString([]byte(secret))
}

func WithJWTAuth(handler http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, ok := utils.GetTokenFromRequest(r)
		if !ok {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no login token provided"))
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
		username := claims["username"].(string)
		userLoginToken := utils.GetLoginTokenFromRequest(r)

		loggedIn, err := store.HasLoginToken(username, userLoginToken)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, err)
			return
		}

		if !loggedIn {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid login token"))
			return
		}
		
		user, err := store.GetUserByUserName(username)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		ctx := r.Context()
		ctx = utils.NewContextWithUser(ctx, user)
		r = r.WithContext(ctx)

		handler(w, r)
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method '%v'", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("permission denied"))
}
