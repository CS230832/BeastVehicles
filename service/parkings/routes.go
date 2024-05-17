package parkings

import (
	"CS230832/BeastVehicles/service/auth"
	"CS230832/BeastVehicles/types"
	"CS230832/BeastVehicles/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ParkingStore
}

func NewHandler(store types.ParkingStore) *Handler {
	return &Handler {store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router, store types.UserStore) {
	router.HandleFunc("/parkings/register", auth.WithJWTAuth(h.registerParking, store)).Methods(http.MethodPost)
	router.HandleFunc("/parkings/delete", auth.WithJWTAuth(h.removeParking, store)).Methods(http.MethodDelete)
	router.HandleFunc("/parkings/info", h.getParking).Methods(http.MethodGet)
}

func (h *Handler) registerParking(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	if user.Role != types.Root {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("only root can register a parking"))
		return
	}

	var payload types.ParkingPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("parking body not given"))
		return
	}

	if err := h.store.AddParking(&payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) removeParking(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	if user.Role != types.Root {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("only root user can delete a parking"))
		return
	}

	if !r.URL.Query().Has("name") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("name parameter not given"))
		return
	}

	parking := r.URL.Query().Get("name")

	if err := h.store.RemoveParking(parking); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) getParking(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("name") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("name parameter not given"))
		return
	}

	name := r.URL.Query().Get("name")

	parking, err := h.store.GetParking(name)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, parking)
}
