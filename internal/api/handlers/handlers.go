package handlers

import "github.com/franciscofferraz/go-struct/internal/db/repositories"

type Handlers struct {
	UserRepository *repositories.UserRepository
	UserHandler    *UserHandler
}

func NewHandler(ur *repositories.UserRepository) *Handlers {
	return &Handlers{
		UserRepository: ur,
		UserHandler:    NewUserHandler(ur),
	}
}
