package parking

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
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router, store types.AdminStore) {
	router.HandleFunc("/parking", auth.WithJWTAuth(h.addParking, store)).Methods(http.MethodPost)
	router.HandleFunc("/parking/{name}/free", h.getAllFreeSlots).Methods(http.MethodGet)
	router.HandleFunc("/parking/{name}/full", h.getAllFullSlots).Methods(http.MethodGet)
}

func (h *Handler) addParking(w http.ResponseWriter, r *http.Request) {
	admin, ok := utils.FromContext(r.Context())

	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("only super admin can add a parking"))
		return
	}
	if !admin.IsSuper {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("only super admin can add a parking"))
		return
	}

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

func (h *Handler) getAllFreeSlots(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	slots, err := h.store.GetAllFreeSlots(name)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var freeSlots map[string][]types.EmptySlotReturnPayload = make(map[string][]types.EmptySlotReturnPayload)

	for _, slot := range slots {
		freeSlots[slot.BlockName] = append(
			freeSlots[slot.BlockName],
			types.EmptySlotReturnPayload{SlotNumber: slot.SlotNumber},
		)
	}

	utils.WriteJSON(w, http.StatusOK, freeSlots)
}

func (h *Handler) getAllFullSlots(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	slots, err := h.store.GetAllFullSlots(name)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	var fullSlots map[string][]types.FullSlotReturnPayload = make(map[string][]types.FullSlotReturnPayload)

	for _, slot := range slots {
		fullSlots[slot.BlockName] = append(
			fullSlots[slot.BlockName],
			types.FullSlotReturnPayload{SlotNumber: slot.SlotNumber, WinCode: slot.WinCode},
		)
	}

	utils.WriteJSON(w, http.StatusOK, fullSlots)
}
