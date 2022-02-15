package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	project "github.com/jitin07/qastack/internal/project"
	release "github.com/jitin07/qastack/internal/release"
	database "github.com/jitin07/qastack/internal/repository"
	transportHttp "github.com/jitin07/qastack/internal/transport/http"

	"net/http"

	log "github.com/sirupsen/logrus"
)

// App- contains application information
type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(
		log.Fields{
			"AppName":    app.Name,
			"AppVersion": app.Version,
		}).Info("Setting up our QAStack App")
	fmt.Println("Setting up our QAStack App")

	var err error
	var db *gorm.DB

	db, err = database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.Migration(db)

	if err != nil {
		return err
	}

	projectService := project.NewService(db)
	releaseService := release.NewService(db)

	fmt.Println(db == nil)
	defer db.Close()

	handler := transportHttp.NewHandler(projectService, releaseService)
	handler.SetupRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"X-Requested-With", "Content-Type", "Authorization", "Referer"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST", "OPTIONS"},
	})

	handlers := c.Handler(handler.Router)

	if err := http.ListenAndServe(":8091", handlers); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

// Our main entrypoint for the application
func main() {
	app := App{
		Name:    "QaStack-Service",
		Version: "1.0.0",
	}
	if err := app.Run(); err != nil {
		fmt.Println(err)
		fmt.Println("Error starting up our REST API")
	}
}
