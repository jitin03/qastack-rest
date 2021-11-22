package response

import (
	"encoding/json"
	"net/http"
)
func setupCorsResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods",  "GET,POST,OPTIONS,PUT,DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

// JSON returns a well formated response with a status code
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	setupCorsResponse(&w)
	
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    
    err := json.NewEncoder(w).Encode(data)
    if err != nil {
		panic(err)
    }
}

// ERROR returns a jsonified error response along with a status code.
func ERROR(w http.ResponseWriter, statusCode int, err error) {
    w.Header().Set("Content-Type", "application/json")
	setupCorsResponse(&w)
    if err != nil {
        JSON(w, statusCode, struct {
            Error string `json:"error"`
        }{
            Error: err.Error(),
        })
        return
    }
    JSON(w, http.StatusBadRequest, nil)
}

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	setupCorsResponse(&w)
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}