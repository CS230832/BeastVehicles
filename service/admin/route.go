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
	defaultAdmin := types.AdminPayload {
		Email: "admin@gmail.com",
		Password: "admin",
		FirstName: "admin",
		LastName: "admin",
		IsSuper: true,
	}
	h.store.AddAdmin(&defaultAdmin)

	router.HandleFunc("/admin/register", auth.WithJWTAuth(h.registerAdmin, h.store)).Methods(http.MethodPost)
	router.HandleFunc("/admin/{email}", auth.WithJWTAuth(h.removeAdmin, h.store)).Methods(http.MethodDelete)
	router.HandleFunc("/admin/login", h.loginAdmin).Methods(http.MethodPost)
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
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if !auth.ComparePasswords(loginPayload.Password, admin.Password) {
		utils.WriteError(w, http.StatusBadGateway, err)
		return
	}

	token, err := auth.NewJWT(config.Envs.JWTSecret, admin.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
}
