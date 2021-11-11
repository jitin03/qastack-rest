package http

import (
	"encoding/json"
	"fmt"
	database "github.com/jitin07/qastack/internal/repository"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jitin07/qastack/internal/project"
	"github.com/jitin07/qastack/internal/release"
	"github.com/jitin07/qastack/internal/transport/response"
	log "github.com/sirupsen/logrus"
)


type Handler struct{
	Router *mux.Router
	Service *project.Service
	ReleaseService *release.Service
}


func NewHandler(service *project.Service,releaseService *release.Service) * Handler{
	return &Handler{
		Service: service,
		ReleaseService: releaseService,
	}
}


func logginMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter,r * http.Request){
		log.WithFields(
			log.Fields{
				"Method":      r.Method,
				"Path":        r.URL.Path,
			}).
			Info("handled request")
		next.ServeHTTP(w, r)
	})
}

func(h *Handler) SetupRouter(){
	log.Info("Setting up Routes")

	h.Router = mux.NewRouter()

	am := AuthMiddleware{database.NewAuthRepository()}
	h.Router.Use(am.authorizationHandler())



	h.Router.HandleFunc("/api/project/{id}",h.GetProject).Methods("GET").Name("GetAProject")
	h.Router.HandleFunc("/api/projects",h.GetAllProjects).Methods("GET").Name("GetAllProjects")
	h.Router.HandleFunc("/api/project/create",h.AddProject).Methods("POST").Name(("NewProject"))
	h.Router.HandleFunc("/api/project/update/{id}",h.UpdateProject).Methods("PUT").Name("UpdateProject")
	h.Router.HandleFunc("/api/project/delete/{id}",h.DeleteProject).Methods("DELETE").Name("DeleteProject")

	h.Router.HandleFunc("/api/release/create",h.AddRelease).Methods("POST").Name("NewRelease")
	h.Router.HandleFunc("/api/releases",h.GetAllRelease).Methods("GET").Name("GetAllRelease")
	h.Router.HandleFunc("/api/release/{id}",h.GetRelease).Methods("GET").Name("GetRelease")
	h.Router.HandleFunc("/api/release/update/{id}",h.UpdateRelease).Methods("PUT").Name("UpdateRelease")

	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter,r *http.Request)  {
		json.NewEncoder(w).Encode("Running...")

	})

}




func(h *Handler)GetRelease(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	release,err :=h.ReleaseService.GetRelease(uint(id))
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w,http.StatusOK,release)
}



func(h * Handler)UpdateRelease(w http.ResponseWriter,r * http.Request){
	var newRelease release.Release
	params := mux.Vars(r)
	// convert the id type from string to int
	releaseId, err := strconv.Atoi(params["id"])

	if err!=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(releaseId)
	reqBody, _ := ioutil.ReadAll(r.Body)

    json.Unmarshal(reqBody, &newRelease)
	release,err :=h.ReleaseService.UpdateRelease(uint(releaseId),newRelease)
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w,http.StatusOK,release)


}

func(h *Handler)AddRelease(w http.ResponseWriter,r *http.Request){

	reqBody, _ :=ioutil.ReadAll(r.Body)
	var newRelease release.Release

	json.Unmarshal(reqBody,&newRelease)

	release,err :=h.ReleaseService.AddRelease(newRelease)

	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w,http.StatusOK,release)

}

func(h *Handler)GetAllRelease(w http.ResponseWriter,r *http.Request){

	release,err :=h.ReleaseService.GetAllRelease()
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w,http.StatusOK,release)

}

