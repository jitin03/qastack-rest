package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jitin07/qastack/internal/project"
	"github.com/jitin07/qastack/internal/transport/response"
)

// Getproject handler- retriever by id integration of project servie + database object
func(h *Handler) GetProject(w http.ResponseWriter, r *http.Request){

	response.SetupCorsResponse(&w,r)
	params := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	project,err := h.Service.GetProject(uint(id))
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w,http.StatusOK,project)
}


func(h * Handler)DeleteProject(w http.ResponseWriter,r *http.Request){
	response.SetupCorsResponse(&w,r)

	params := mux.Vars(r)
	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	error :=h.Service.DeleteProject(uint(id))
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, error)
		return
	}
	
	response.JSON(w,http.StatusOK,"success")
}

func (h *Handler) GetAllProjects(w http.ResponseWriter, r *http.Request){
	response.SetupCorsResponse(&w,r)

	project,err :=h.Service.GetAllProjects()
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w,http.StatusOK,project)
}


func(h * Handler) AddProject(w http.ResponseWriter, r * http.Request){
	response.SetupCorsResponse(&w,r)
	// get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(r.Body)
	var newProject project.Project 
    json.Unmarshal(reqBody, &newProject)
	
	project,err :=h.Service.AddProject(newProject)
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w,http.StatusOK,project)
}


func(h * Handler) UpdateProject(w http.ResponseWriter, r * http.Request){
	response.SetupCorsResponse(&w,r)
	params := mux.Vars(r)


	// convert the id type from string to int
	projectId, err := strconv.Atoi(params["id"])
	if err!=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(projectId)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newProject project.Project 
    json.Unmarshal(reqBody, &newProject)
	project,err :=h.Service.UpdateProject(uint(projectId),newProject)
	if err !=nil{
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w,http.StatusOK,project)
}