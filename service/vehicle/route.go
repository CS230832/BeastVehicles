package vehicle

import (
	"CS230832/BeastVehicles/service/auth"
	"CS230832/BeastVehicles/types"
	"CS230832/BeastVehicles/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.VehicleStore
}

func NewHandler(store types.VehicleStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router, store types.AdminStore) {
	router.HandleFunc("/vehicle", auth.WithJWTAuth(h.addVehicle, store)).Methods(http.MethodPost)
	router.HandleFunc("/vehicle/{wincode}", auth.WithJWTAuth(h.removeVehicle, store)).Methods(http.MethodDelete)
	router.HandleFunc("/vehicle/{wincode}", h.getVehicle).Methods(http.MethodGet)
}

func (h *Handler) addVehicle(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.FromContext(r.Context())

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("only admins cand add a vehicle"))
		return
	}

	var payload types.VehiclePayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", err.Error()))
		return
	}

	result, err := h.store.AddVehicle(&payload)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}

func (h *Handler) removeVehicle(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.FromContext(r.Context())

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("only admins cand add a vehicle"))
		return
	}

	wincode := mux.Vars(r)["wincode"]

	if len(wincode) == 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wincode cannot be empty"))
		return
	}

	result, err := h.store.RemoveVehicle(wincode)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}

func (h *Handler) getVehicle(w http.ResponseWriter, r *http.Request) {
	wincode := mux.Vars(r)["wincode"]

	result, err := h.store.GetVehicle(wincode)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}