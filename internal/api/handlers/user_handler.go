package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/franciscofferraz/go-struct/internal/api/models"
	"github.com/franciscofferraz/go-struct/internal/customerrors"
	"github.com/franciscofferraz/go-struct/internal/db/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepository *repositories.UserRepository
}

func NewUserHandler(ur *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: ur,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		customerrors.ErrorResponse(w, r, http.StatusBadRequest, "invalid request payload")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(user.Password, 14)
	if err != nil {
		customerrors.ErrorResponse(w, r, http.StatusInternalServerError, "password hashing failed")
		return
	}

	user.Password = hashedPassword

	if err := uh.UserRepository.Create(&user); err != nil {
		customerrors.ServerErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
