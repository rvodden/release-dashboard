package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"release-dashboard/app/model"
	"strconv"
)

type ApiController struct {
}

func NewApiController() *ApiController {
	return new(ApiController)
}

func (apiController *ApiController) Register(router *mux.Router, path string) {
	log.Println("Registering API Router")
	subRouter := router.PathPrefix(path).Subrouter()
	subRouter.HandleFunc("/releaseTrains", getReleaseTrains).Methods(http.MethodGet)
	subRouter.HandleFunc("/releaseTrains/{id}", getReleaseTrain).Methods(http.MethodGet)
}

func getReleaseTrains(writer http.ResponseWriter, request *http.Request) {
	releaseTrains := []model.ReleaseTrain{{ID: 1, Name: "Mock Release Train"}}
	json.NewEncoder(writer).Encode(releaseTrains)
}

func getReleaseTrain(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Println(err)
		http.Error(writer, "ID must be a valid integer ("+err.Error()+")", http.StatusBadRequest)
	} else {
		releaseTrain := model.ReleaseTrain{ID: id, Name: "Mock Release Train"}
		json.NewEncoder(writer).Encode(releaseTrain)
	}
}
