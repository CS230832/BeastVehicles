package parking

import (
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
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	router.HandleFunc("/parking", h.AddParking).Methods(http.MethodPost)
	router.HandleFunc("/parking/free", h.GetAllFreeSlots).Methods(http.MethodGet)
	router.HandleFunc("/parking/full", h.GetAllFullSlots).Methods(http.MethodGet)
}

func (h *Handler) AddParking(w http.ResponseWriter, r *http.Request) {
	var payload types.ParkingAddPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", err.Error()))
		return
	}

	if err := h.store.AddParking(&payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, payload)
}

func (h *Handler) GetAllFreeSlots(w http.ResponseWriter, r *http.Request) {
	var payload types.SlotPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", err.Error()))
		return
	}

	slots, err := h.store.GetAllFreeSlots(payload.ParkingName)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, slots)
}

func (h *Handler) GetAllFullSlots(w http.ResponseWriter, r *http.Request) {
	var payload types.SlotPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %s", err.Error()))
		return
	}

	slots, err := h.store.GetAllFullSlots(payload.ParkingName)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, slots)
}
