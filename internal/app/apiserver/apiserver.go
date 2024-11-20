package apiserver

import (
	"io"
	"net/http"

	"github.com/EWinterhalter/go_api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type apiserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *apiserver {
	return &apiserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *apiserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *apiserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)

	return nil
}

func (s *apiserver) configRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *apiserver) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

func (s *apiserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
