package apiserver

import (
	"github.com/sirupsen/logrus"
)

type apiserver struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *apiserver {
	return &apiserver{
		config: config,
		logger: logrus.New(),
	}
}

func (s *apiserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("starting")
	return nil
}

func (s *apiserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}
