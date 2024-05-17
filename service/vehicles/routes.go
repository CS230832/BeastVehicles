package vehicles

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

func (h *Handler) RegisterRoutes(router *mux.Router, store types.UserStore) {
	router.HandleFunc("/vehicles/register", auth.WithJWTAuth(h.addVehicle, store)).Methods(http.MethodPost)
	router.HandleFunc("/vehicles/delete", auth.WithJWTAuth(h.removeVehicle, store)).Methods(http.MethodDelete)
	router.HandleFunc("/vehicles/info", h.getVehicle).Methods(http.MethodGet)

	router.HandleFunc("/vehicles/set/register", auth.WithJWTAuth(h.addVehicleSet, store)).Methods(http.MethodPost)
	router.HandleFunc("/vehicles/set/delete", auth.WithJWTAuth(h.removeVehicleSet, store)).Methods(http.MethodDelete)
	router.HandleFunc("/vehicles/set/info", h.getVehicleSet).Methods(http.MethodPost)
}

func (h *Handler) addVehicle(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	if user.ParkingName == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("only users with parking can register/delete vehicles from parkings"))
		return
	}

	var payload types.VehicleRegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no wincode is given"))
		return
	}

	vehicle, err := h.store.AddVehicle(&payload, user.ParkingName)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, vehicle)
}

func (h *Handler) removeVehicle(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	if user.ParkingName == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("only users with parking can register/delete vehicles from parkings"))
		return
	}

	if !r.URL.Query().Has("wincode") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wincode parameter not given"))
		return
	}

	wincode := r.URL.Query().Get("wincode")

	if err := h.store.RemoveVehicle(wincode); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) getVehicle(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("wincode") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("wincode parameter not given"))
		return
	}

	wincode := r.URL.Query().Get("wincode")

	vehicle, err := h.store.GetVehicle(wincode)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, vehicle)
}

func (h *Handler) addVehicleSet(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	if user.ParkingName == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("only users with parking can register/delete vehicles from parkings"))
		return
	}

	var payload types.VehicleSetRegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no wincode is given"))
		return
	}

	vehicles, err := h.store.AddVehicles(&payload, user.ParkingName)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, vehicles)
}

func (h *Handler) removeVehicleSet(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	if user.ParkingName == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("only users with parking can register/delete vehicles from parkings"))
		return
	}

	var payload types.VehicleSetRegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no wincode is given"))
		return
	}

	if err := h.store.RemoveVehicles(payload.Set); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) getVehicleSet(w http.ResponseWriter, r *http.Request) {
	var payload types.VehicleSetRegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no wincode is given"))
		return
	}

	vehicles, err := h.store.GetVehicles(payload.Set)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, vehicles)
}
