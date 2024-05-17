package users

import (
	"CS230832/BeastVehicles/config"
	"CS230832/BeastVehicles/service/auth"
	"CS230832/BeastVehicles/types"
	"CS230832/BeastVehicles/utils"
	"fmt"
	"log"
	"net/http"

	_ "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	if _, err := h.store.GetUser(config.Envs.RootUserName); err != nil {
		if err := h.store.AddUser(&types.UserRegisterPayload{
			UserName:  config.Envs.RootUserName,
			Password:  config.Envs.RootPassword,
			Role:      types.Root,
			FirstName: config.Envs.RootFirstName,
			LastName:  config.Envs.RootLastName,
		}); err != nil {
			log.Fatalf("failed to add root user: '%s'", err.Error())
		}
	}

	router.HandleFunc("/users/register", auth.WithJWTAuth(h.registerUser, h.store)).Methods(http.MethodPost)
	router.HandleFunc("/users/info", auth.WithJWTAuth(h.GetUser, h.store)).Methods(http.MethodGet)
	router.HandleFunc("/users/delete", auth.WithJWTAuth(h.removeUser, h.store)).Methods(http.MethodDelete)
	router.HandleFunc("/users/login", h.loginUser).Methods(http.MethodPost)
	router.HandleFunc("/users/logout", auth.WithJWTAuth(h.logoutUser, h.store)).Methods(http.MethodPost)
}

func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	if user.Role != types.Root && user.Role != types.Manager {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
		return
	}

	var payload types.UserRegisterPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user body not given"))
		return
	}

	if user.Role == types.Manager && payload.Role == types.Root {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("only root user can add another root"))
		return
	}

	if user.Role == types.Manager && (payload.ParkingName != user.ParkingName) {
		utils.WriteError(w, http.StatusForbidden, fmt.Errorf("manager cannot add user to another parking"))
		return
	}

	if err := h.store.AddUser(&payload); err != nil {
		utils.WriteError(w, http.StatusExpectationFailed, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("username") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no username parameter given"))
		return
	}

	username := r.URL.Query().Get("username")

	user, err := h.store.GetUser(username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, user)
}

func (h *Handler) removeUser(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	if user.Role == types.Admin {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("only root and manager can delete a user"))
		return
	}

	if !r.URL.Query().Has("username") {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("no username parameter given"))
		return
	}

	username := r.URL.Query().Get("username")

	targetUser, err := h.store.GetUser(username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if user.Role == types.Root && user.Role == targetUser.Role && user.UserName != targetUser.UserName {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("root user can delete another root user only if it is themselves"))
		return
	}

	if user.Role == types.Manager && targetUser.Role == types.Root {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("manager cannot delete a root user"))
		return
	}

	if user.Role == types.Manager && user.Role == targetUser.Role && user.UserName != targetUser.UserName {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("manager can delete another manager only if it is themselves"))
		return
	}

	if user.Role == types.Manager && targetUser.ParkingName != user.ParkingName {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("user is not in your parking"))
		return
	}

	if err := h.store.RemoveAllLoginTokens(username); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	if err := h.store.RemoveUser(username); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	var payload types.UserLoginPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user body not given"))
		return
	}

	user, err := h.store.GetUser(payload.UserName)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	if !auth.ComparePasswords(payload.Password, user.Password) {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid credentials"))
		return
	}

	token, err := auth.NewJWT(config.Envs.JWTSecret, user.UserName)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	loginToken := utils.GetLoginTokenFromRequest(r)
	if err := h.store.AddLoginToken(user.UserName, loginToken); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, token)
}

func (h *Handler) logoutUser(w http.ResponseWriter, r *http.Request) {
	user, ok := utils.GetUserFromContext(r.Context())
	if !ok {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get user from context"))
		return
	}

	loginToken := utils.GetLoginTokenFromRequest(r)
	if err := h.store.RemoveLoginToken(user.UserName, loginToken); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}
