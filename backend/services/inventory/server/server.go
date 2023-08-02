package server

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/serhiq/effeInventory/pkg/store/mysql"
	config "github.com/serhiq/effeInventory/services/inventory/configs"
	"log"
)

type Server struct {
	cfg       config.Config
	store     *mysql.Store
	startFunc []func()
	stopFunc  []func()
}

func Serve(cfg config.Config) (*Server, error) {
	fmt.Printf("\n%#v\n\n", cfg)

	var s = &Server{
		cfg:       cfg,
		store:     nil,
		startFunc: nil,
		stopFunc:  nil,
	}

	for _, init := range []func() error{
		s.initDb,
	} {
		if err := init(); err != nil {
			return nil, errors.Wrap(err, "serve failed")
		}
	}
	return s, nil
}

func (s *Server) Start() error {
	fmt.Println("Server is starting...")

	for _, start := range s.startFunc {
		start()
	}

	return nil
}

func (s *Server) Stop() {
	for _, stop := range s.stopFunc {
		stop()
	}
}

func (s *Server) initDb() error {
	store, err := mysql.New(s.cfg.DBConfig)

	if err != nil {
		return err
	}

	s.store = store

	s.addStopDelegate(func() {
		db, err := s.store.Db.DB()
		if err != nil {
			log.Printf("database: error close database, %s", err)
			return
		}
		err = db.Close()
		if err != nil {
			log.Printf("database: error close database, %s", err)
			return
		}
		log.Print("database: close")
	})
	return err
}

func (s *Server) addStartDelegate(delegate func()) {
	s.startFunc = append(s.startFunc, delegate)
}

func (s *Server) addStopDelegate(delegate func()) {
	s.stopFunc = append(s.stopFunc, delegate)
}
