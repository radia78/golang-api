package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radia78/personal-finance-api/types"
	"github.com/radia78/personal-finance-api/utils"
)

type Handler struct {
}

// Create a new handler
func NewHandler() *Handler {
	return &Handler{}
}

// Create a method to register the routes for the user service
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// Create a handler function that handles user login
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

// Create a handler function that handles user registration
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// Check if the user exists
	// If not, we create the user
}
