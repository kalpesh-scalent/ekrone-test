package restAPI

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kalpesh-scalent/ekrone-test/pkg/logger"
	"github.com/kalpesh-scalent/ekrone-test/pkg/service"
)

type APIRouter struct {
	projectService service.ProjectService
	logger         *logger.Logger
}

func GetAPIRouter(ps service.ProjectService, l *logger.Logger) *APIRouter {
	return &APIRouter{
		projectService: ps,
		logger:         l,
	}
}

func (a APIRouter) InitAPIRoutes(r *mux.Router) {

	r = r.PathPrefix("/v1").Subrouter()

	r.HandleFunc("/projects/{num:[0-9]+}", a.GetProjects).Methods("GET")
}

func (a APIRouter) sendResponseData(w http.ResponseWriter, resp interface{}) {
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		log.Println(err)
		http.Error(w, "error in encoding jason", http.StatusInternalServerError)
		return
	}
	w.Header().Add("context-type", "application-json")
}

func (a APIRouter) sendErrorResponse(w http.ResponseWriter, statusCode int, errorMessage string) {

	resp := map[string]string{"error": errorMessage}
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		log.Println(err)
		http.Error(w, "error in encoding jason", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Header().Add("context-type", "application-json")
}
