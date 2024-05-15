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

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if v == nil {
		return nil
	}
	// return json.NewEncoder(w).Encode(model{Success: true, Data: v})
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(model{Success: false, Data: err.Error()}); err != nil {
		log.Println(err.Error())
	}
}

func GetTokenFromRequest(r *http.Request) (string, error) {
	if token := r.Header.Get("Authorization"); token != "" {
		return token, nil
	}
	
	if r.URL.Query().Has("token") {
		return r.URL.Query().Get("token"), nil
	}

	return "", fmt.Errorf("should give jwt token")
}

type key int
var adminKey key

func NewContext(ctx context.Context, admin *types.AdminPayload) context.Context {
	return context.WithValue(ctx, adminKey, admin)
}

func FromContext(ctx context.Context) (*types.AdminPayload, bool) {
	admin, ok := ctx.Value(adminKey).(*types.AdminPayload)
	return admin, ok
}