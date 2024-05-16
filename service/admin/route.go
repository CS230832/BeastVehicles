package admin

import (
	"CS230832/BeastVehicles/config"
	"CS230832/BeastVehicles/service/auth"
	"CS230832/BeastVehicles/types"
	"CS230832/BeastVehicles/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.AdminStore
}

func NewHandler(store types.AdminStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	if _, err := h.store.GetAdmin("admin@gmail.com"); err != nil {
		defaultAdmin := types.AdminPayload{
			Email:     config.Envs.AdminEmail,
			Password:  config.Envs.AdminPassword,
			FirstName: config.Envs.AdminFirstName,
			LastName:  config.Envs.AdminLastName,
			IsSuper:   true,
		}
		h.store.AddAdmin(&defaultAdmin)
	}

	router.HandleFunc("/admin/register", auth.WithJWTAuth(h.registerAdmin, h.store)).Methods(http.MethodPost)
	router.HandleFunc("/admin/{email}", auth.WithJWTAuth(h.removeAdmin, h.store)).Methods(http.MethodDelete)
	router.HandleFunc("/admin/login", h.loginAdmin).Methods(http.MethodPost)
	router.HandleFunc("/admin/logout", auth.WithJWTAuth(h.logoutAdmin, h.store)).Methods(http.MethodPost)
	router.HandleFunc("/admin/info", auth.WithJWTAuth(h.getAdminInfo, h.store)).Methods(http.MethodGet)
}

func (h *Handler) registerAdmin(w http.ResponseWriter, r *http.Request) {
	admin, ok := utils.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("admin should log in"))
		return
	}

	if !admin.IsSuper {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("admin should be super"))
		return
	}

	var payload types.AdminPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.store.AddAdmin(&payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) removeAdmin(w http.ResponseWriter, r *http.Request) {
	admin, ok := utils.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("should log in"))
		return
	}

	email := mux.Vars(r)["email"]

	if !admin.IsSuper {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("should be super admin"))
		return
	}

	if err := h.store.RemoveAdmin(email); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) loginAdmin(w http.ResponseWriter, r *http.Request) {
	var loginPayload types.AdminLoginPayload
	if err := utils.ParseJSON(r, &loginPayload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	admin, err := h.store.GetAdmin(loginPayload.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("incorrenct password or email"))
		return
	}

	if !auth.ComparePasswords(loginPayload.Password, admin.Password) {
		utils.WriteError(w, http.StatusBadGateway, fmt.Errorf("incorrenct password or email"))
		return
	}

	token, err := auth.NewJWT(config.Envs.JWTSecret, admin.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	userAgent := r.Header.Get("User-Agent")

	if userAgent == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user-agent should be passed"))
		return
	}

	h.store.AddToken(admin.Email, userAgent)

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}

func (h *Handler) logoutAdmin(w http.ResponseWriter, r *http.Request) {
	admin, ok := utils.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("should log in"))
		return
	}

	if err := h.store.RemoveToken(admin.Email, r.Header.Get("User-Agent")); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) getAdminInfo(w http.ResponseWriter, r *http.Request) {
	admin, ok := utils.FromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("should log in"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		IsSuper   bool   `json:"is_super"`
	}{
		FirstName: admin.FirstName,
		LastName:  admin.LastName,
		IsSuper:   admin.IsSuper,
	})
}
