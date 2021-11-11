package http

import (
	"github.com/gorilla/mux"
	"github.com/jitin07/qastack/internal/errs"
	database "github.com/jitin07/qastack/internal/repository"
	"github.com/jitin07/qastack/internal/transport/response"
	_ "github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	repo database.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)
				log.Info("token",token)
				log.Info(currentRoute.GetName())
				log.Info(currentRouteVars)
				isAuthorized := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)
				log.Info("authorised:",isAuthorized)
				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					appError := errs.AppError{http.StatusForbidden, "Unauthorized"}
					response.WriteResponse(w, appError.Code, appError.AsMessage())
				}
			} else {
				response.WriteResponse(w, http.StatusUnauthorized, "missing token")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	/*
	   token is coming in the format as below
	   "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50cyI6W.yI5NTQ3MCIsIjk1NDcyIiw"
	*/
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
