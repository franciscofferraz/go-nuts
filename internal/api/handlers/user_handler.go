package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/franciscofferraz/go-struct/internal/api/models"
	"github.com/franciscofferraz/go-struct/internal/db/repositories"
	"github.com/franciscofferraz/go-struct/internal/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	*Handlers
}

func NewUserHandler(userRepo *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		Handlers: NewHandler(userRepo),
	}
}

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errors.ErrorResponse(w, r, http.StatusBadRequest, "invalid request payload")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(user.Password, 14)
	if err != nil {
		errors.ErrorResponse(w, r, http.StatusInternalServerError, "password hashing failed")
		return
	}

	user.Password = hashedPassword

	if err := uh.UserRepository.Create(&user); err != nil {
		errors.ServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
