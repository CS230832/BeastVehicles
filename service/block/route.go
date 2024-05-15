package block

import (
	"CS230832/BeastVehicles/types"
	"CS230832/BeastVehicles/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	store types.BlockStore
}

func NewHandler(store types.BlockStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/block/free", h.getFreeSlots).Methods(http.MethodGet)
	router.HandleFunc("/block/full", h.getFullSlots).Methods(http.MethodGet)
}

func (h *Handler) getFreeSlots(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	if !queryParams.Has("parking") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("'parking' query param should be given"))
		return
	}
	if !queryParams.Has("name") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("'name' query param should be given"))
		return
	}

	parking := queryParams.Get("parking")
	name := queryParams.Get("name")

	slots, err := h.store.GetFreeSlots(parking, name)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	var freeSlots []types.EmptySlotReturnPayload

	for _, slot := range slots {
		freeSlots = append(freeSlots, types.EmptySlotReturnPayload{SlotNumber: slot.SlotNumber})
	}

	utils.WriteJSON(w, http.StatusOK, freeSlots)
}

func (h *Handler) getFullSlots(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	if !queryParams.Has("parking") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("'parking' query param should be given"))
		return
	}
	if !queryParams.Has("name") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("'name' query param should be given"))
		return
	}

	parking := queryParams.Get("parking")
	name := queryParams.Get("name")

	slots, err := h.store.GetFullSlots(parking, name)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("internal server error"))
		return
	}

	var fullSlots []types.FullSlotReturnPayload

	for _, slot := range slots {
		fullSlots = append(fullSlots, types.FullSlotReturnPayload{SlotNumber: slot.SlotNumber, WinCode: slot.WinCode})
	}

	utils.WriteJSON(w, http.StatusOK, fullSlots)
}
