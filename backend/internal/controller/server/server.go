package server

import (
	"io"
	"net/http"

	"github.com/HackYourCareer/SmartKickers/internal/model"

	"github.com/gorilla/mux"
)

type Server interface {
	Start() error
	createResponse(io.Reader) ([]byte, error)
}
type server struct {
	router  *mux.Router
	address string
	game    model.Game
}

func New(addr string, game model.Game) Server {
	serv := server{
		router:  mux.NewRouter(),
		address: addr,
		game:    game,
	}
	serv.router.HandleFunc("/", serv.TableMessagesHandler)
	serv.router.HandleFunc("/reset", serv.ResetScoreHandler).Methods("PUT")

	return serv
}

func (s server) Start() error {
	return http.ListenAndServe(s.address, s.router)
}
