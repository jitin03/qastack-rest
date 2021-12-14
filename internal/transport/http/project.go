package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/jitin07/qastack/internal/project"
	"github.com/jitin07/qastack/internal/transport/response"
)

// Getproject handler- retriever by id integration of project servie + database object
func(h *Handler) GetProject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	if r.Method == http.MethodOptions {

		return
	}
	

	params := mux.Vars(r)
	// convert the id type from string to int
	// id, err := strconv.Atoi(params["id"])
	project,err := h.Service.GetProject(params["id"])
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w,http.StatusOK,project)
}


func(h * Handler)DeleteProject(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	if r.Method == http.MethodOptions {

		return
	}

	params := mux.Vars(r)
	// convert the id type from string to int
	// id, err := strconv.Atoi(params["id"])

	error :=h.Service.DeleteProject(params["id"])
	if error !=nil{
		response.ERROR(w, http.StatusInternalServerError, error)
		return
	}
	
	response.JSON(w,http.StatusOK,"success")
}

func (h *Handler) GetAllProjects(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	if r.Method == http.MethodOptions {
		return
	}
	query := r.URL.Query()
	filters := query.Get("userid")
	id, err := strconv.Atoi(filters)
	project,err :=h.Service.GetAllProjects(id)
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w,http.StatusOK,project)
}


func(h * Handler) AddProject(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	log.Info(r.Method)
	if r.Method == http.MethodOptions {

		 response.JSON(w,200,"true")
	}
	// get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(r.Body)
	var newProject project.Project
	log.Info(reqBody)

    json.Unmarshal(reqBody, &newProject)


	project,err :=h.Service.AddProject(newProject)
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w,200,project)
}


func(h * Handler) UpdateProject(w http.ResponseWriter, r * http.Request){

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	params := mux.Vars(r)
	if r.Method == http.MethodOptions {

		response.JSON(w,200,"true")
	}

	// convert the id type from string to int
	// projectId, err := strconv.Atoi(params["id"])
	// if err!=nil{
	// 	response.ERROR(w, http.StatusInternalServerError, err)
	// 	return
	// }
	// fmt.Println(projectId)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newProject project.Project 
    json.Unmarshal(reqBody, &newProject)
	project,err :=h.Service.UpdateProject(params["id"],newProject)
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w,http.StatusOK,project)
}