package blocks

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
	router.HandleFunc("/blocks/info", h.getBlock).Methods(http.MethodGet)
	router.HandleFunc("/blocks/free", h.getBlockFreeSlots).Methods(http.MethodGet)
	router.HandleFunc("/blocks/full", h.getBlockFullSlots).Methods(http.MethodGet)

	router.HandleFunc("/blocks/all/info", h.getBlocks).Methods(http.MethodGet)
	router.HandleFunc("/blocks/all/free", h.getFreeSlotsInBlocks).Methods(http.MethodGet)
	router.HandleFunc("/blocks/all/full", h.getFullSlotsInBlocks).Methods(http.MethodGet)
}

func (h *Handler) getBlock(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	if !queries.Has("name") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("name parameter not given"))
		return
	}
	if !queries.Has("parking") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("parking parameter not given"))
		return
	}

	name := queries.Get("name")
	parking := queries.Get("parking")

	block, err := h.store.GetBlock(name, parking)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]types.SlotPayload{block.Name: block.Slots})
}

func (h *Handler) getBlockFreeSlots(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	if !queries.Has("name") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("name parameter not given"))
		return
	}
	if !queries.Has("parking") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("parking parameter not given"))
		return
	}

	name := queries.Get("name")
	parking := queries.Get("parking")

	block, err := h.store.GetBlockFreeSlots(name, parking)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]types.SlotPayload{block.Name: block.Slots})
}

func (h *Handler) getBlockFullSlots(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	if !queries.Has("name") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("name parameter not given"))
		return
	}
	if !queries.Has("parking") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("parking parameter not given"))
		return
	}

	name := queries.Get("name")
	parking := queries.Get("parking")

	block, err := h.store.GetBlockFullSlots(name, parking)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string][]types.SlotPayload{block.Name: block.Slots})
}

func (h *Handler) getBlocks(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	if !queries.Has("parking") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("parking parameter not given"))
		return
	}

	parking := queries.Get("parking")

	blocks, err := h.store.GetBlocks(parking)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, blocks.Blocks)
}

func (h *Handler) getFreeSlotsInBlocks(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	if !queries.Has("parking") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("parking parameter not given"))
		return
	}

	parking := queries.Get("parking")

	blocks, err := h.store.GetFreeSlotsInBlocks(parking)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, blocks.Blocks)
}

func (h *Handler) getFullSlotsInBlocks(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()

	if !queries.Has("parking") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("parking parameter not given"))
		return
	}

	parking := queries.Get("parking")

	blocks, err := h.store.GetFullSlotsInBlocks(parking)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, blocks.Blocks)
}
