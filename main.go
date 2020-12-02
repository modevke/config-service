package main

import (
	"fmt"
	"log"
	rn "math/rand"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"github.com/modevke/config-service/docs"
	"github.com/modevke/config-service/infrastructure/utility"
	"github.com/modevke/config-service/interfaces"
)

func init() {
	rn.Seed(time.Now().UnixNano())
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// @title Configuration Service API
// @version 1.0.0
// @description Configuration service API.
func main() {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	// LOAD AND VALIDATE ENV VARIABLES
	var host string
	var envVars utility.EnvironmentVariables
	if logger := godotenv.Load(); logger != nil {
		err = fmt.Errorf("enf file is missing")
		return
	}
	envVars.Port = os.Getenv("API_PORT")
	envVars.Host = os.Getenv("API_HOST")
	envVars.ProjectName = os.Getenv("PROJECT_NAME")
	envVars.Scheme = os.Getenv("API_SCHEME")
	envVars.Environment = os.Getenv("GO_ENV")

	if verr := envVars.Validate(); verr != nil {
		err = verr
		return
	}

	// SWAGGER VARIABLES
	host = envVars.Host + ":" + envVars.Port
	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Schemes = []string{envVars.Scheme}
	docs.SwaggerInfo.BasePath = "/api/v1"

	// ROUTER
	router := interfaces.Routing(&envVars)
	app := http.Server{
		Handler: router,
		Addr:    host,

		ReadTimeout:       1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	if serr := app.ListenAndServe(); serr != http.ErrServerClosed{
		err = fmt.Errorf("Unable to start server: %v", serr)
		return
	}

}
