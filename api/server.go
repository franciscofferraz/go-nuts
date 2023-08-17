package api

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/franciscofferraz/go-struct/internal/config"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type Server struct {
	router *httprouter.Router
	db     *sql.DB
	logger *zap.SugaredLogger
	wg     *sync.WaitGroup
	config *config.Config
}

func NewServer(logger *zap.SugaredLogger, db *sql.DB, config *config.Config) *Server {
	server := &Server{
		router: httprouter.New(),
		db:     db,
		logger: logger,
		config: config,
	}

	server.routes()

	return server
}

func (s *Server) Start(addr string) error {
	srv := &http.Server{
		Addr:         addr,
		Handler:      s.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sign := <-quit

		s.logger.Infow("caught signal", "signal", sign.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		s.logger.Infow("completing background tasks", "addr", srv.Addr)

		s.wg.Wait()
		shutdownError <- nil
	}()

	s.logger.Infow("starting server", "addr", srv.Addr, "env", s.config.Env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	s.logger.Infow("stopped server", "addr", srv.Addr)

	return nil
}
