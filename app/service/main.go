package service

import (
	"strconv"

	"github.com/sergionunezgo/gorest/app/service/webserver"
	"github.com/sergionunezgo/gorest/internal/logger"
)

type Config struct {
	Port     int
	Host     string
	LogLevel string
}

type Service interface {
	Start() error
	Close() error
}

func New(cfg *Config) Service {
	logger.Log.Info("initializing api service")
	return webserver.NewHttpService(cfg.Host, strconv.Itoa(cfg.Port))
}
