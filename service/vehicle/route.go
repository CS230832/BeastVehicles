package vehicle

import (
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

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/vehicle", h.AddVehicle).Methods(http.MethodPost)
	router.HandleFunc("/vehicle/{wincode}", h.RemoveVehicle).Methods(http.MethodDelete)
	router.HandleFunc("/vehicle/{wincode}", h.GetVehicle).Methods(http.MethodGet)
}

func (h *Handler) AddVehicle(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) RemoveVehicle(w http.ResponseWriter, r *http.Request) {
	wincode := mux.Vars(r)["wincode"]

	result, err := h.store.RemoveVehicle(wincode)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}

func (h *Handler) GetVehicle(w http.ResponseWriter, r *http.Request) {
	wincode := mux.Vars(r)["wincode"]

	result, err := h.store.GetVehicle(wincode)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, result)
}