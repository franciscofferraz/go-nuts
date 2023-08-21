package config

import (
	"database/sql"
	"flag"
	"os"

	"github.com/franciscofferraz/go-struct/internal/api/handlers"
	"github.com/franciscofferraz/go-struct/internal/db/repositories"
)

type Config struct {
	Port string
	Env  string
	DB   struct {
		DSN          string
		MaxOpenConns int
		MaxIdleConns int
		MaxIdleTime  string
	}
	Limiter struct {
		RPS     float64
		Burst   int
		Enabled bool
	}
	SMTP struct {
		Host     string
		Port     int
		Username string
		Password string
		Sender   string
	}
}

func NewConfig() *Config {
	return &Config{}
}

func (cfg *Config) ParseFlags() error {
	flag.StringVar(&cfg.Port, "port", os.Getenv("PORT"), "API server port")
	flag.StringVar(&cfg.Env, "env", os.Getenv("ENV"), "Environment (development|staging|production)")

	flag.StringVar(&cfg.DB.DSN, "db-dsn", os.Getenv("DSN"), "PostgreSQL DSN")

	flag.IntVar(&cfg.DB.MaxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.DB.MaxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.StringVar(&cfg.DB.MaxIdleTime, "db-max-idle-time", "15m", "PostgreSQL max connection idle time")

	flag.Float64Var(&cfg.Limiter.RPS, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.Limiter.Burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.Limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.StringVar(&cfg.SMTP.Host, "smtp-host", os.Getenv("SMTP_HOST"), "SMTP host")
	flag.IntVar(&cfg.SMTP.Port, "smtp-port", 25, os.Getenv("SMTP_PORT"))
	flag.StringVar(&cfg.SMTP.Username, "smtp-username", os.Getenv("SMTP_USERNAME"), "SMTP username")
	flag.StringVar(&cfg.SMTP.Password, "smtp-password", os.Getenv("SMTP_PASSWORD"), "SMTP password")
	flag.StringVar(&cfg.SMTP.Sender, "smtp-sender", "openMovie <no-reply@test.user.net>", "SMTP sender")

	return nil
}

func (c *Config) InitializeHandlers(db *sql.DB) (*handlers.Handlers, error) {
	return handlers.NewHandler(repositories.NewUserRepository(db)), nil
}
