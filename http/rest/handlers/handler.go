package handlers

import (
	toDoRepo "rest-api/internal/todo/repository"
	toDoService "rest-api/internal/todo/service"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger *logrus.Logger
	router *mux.Router
	toDoService toDoService.Service
}

func newHanadler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger: lg,
		toDoService: toDoService.NewService(toDoRepo.NewRepository(db)),
	}
}