package controller

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"release-dashboard/app/client"
)

type ClientController struct {
}

func NewClientController() *ClientController {
	return new(ClientController)
}

func (clientController *ClientController) Register(router *mux.Router, path string) {
	log.Println("Registering Client Router")
	subRouter := router.PathPrefix(path).Subrouter()
	subRouter.PathPrefix("").Handler(http.StripPrefix(path, http.FileServer(neuteredFileSystem{client.Client})))
}
