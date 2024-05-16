package utils

import (
	"CS230832/BeastVehicles/types"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type model struct {
	Success bool `json:"status"`
	Data    any  `json:"data,omitempty"`
}

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(model{Success: true, Data: value})
}

func WriteError(w http.ResponseWriter, status int, msg error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if msg == nil {
		msg = fmt.Errorf("unknown error")
	}
	if err := json.NewEncoder(w).Encode(model{Success: false, Data: msg.Error()}); err != nil {
		log.Println(err.Error())
	}
}

func GetTokenFromRequest(r *http.Request) (string, bool) {
	if token := r.Header.Get("Authorization"); token != "" {
		return token, true
	}
	
	if r.URL.Query().Has("token") {
		return r.URL.Query().Get("token"), true
	}

	cookie, err := r.Cookie("jwt_token")
	if err == nil {
		return cookie.Value, true
	}

	return "", false
}

type key int
var userKey key

func NewContextWithUser(ctx context.Context, user *types.UserPayload) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func GetUserFromContext(ctx context.Context) (*types.UserPayload, bool) {
	user, ok := ctx.Value(userKey).(*types.UserPayload)
	return user, ok
}

func GetLoginTokenFromRequest(r *http.Request) string {
	return fmt.Sprintf("%s, %s", r.Host, r.Header.Get("User-Agent"))
}
