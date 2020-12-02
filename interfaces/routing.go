package interfaces

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/modevke/config-service/infrastructure/utility"
)

func Routing(envVars *utility.EnvironmentVariables) http.Handler {

	r := mux.NewRouter()

	// ROUTER GROUPS
	api := r.PathPrefix("/api/v1").Subrouter()

	// Country APIs
	countryRoutes := api.PathPrefix("/country").Subrouter()
	country := CountryHandler{
		Environment: envVars.Environment,
	}


	countryRoutes.HandleFunc("/create", country.CreateCountry).Methods("POST")
	countryRoutes.HandleFunc("/fetch-id/{id}", country.FetchCountryByID).Methods("GET")
	countryRoutes.HandleFunc("/fetch-iso/{iso}", country.FetchCountryByISO).Methods("GET")
	countryRoutes.HandleFunc("/fetch-all", country.FetchCountries).Methods("GET")
	countryRoutes.HandleFunc("/update", country.UpdateCountry).Methods("PUT")
	countryRoutes.HandleFunc("/delete-soft", country.SoftDeleteCountry).Methods("DELETE")
	countryRoutes.HandleFunc("/delete-hard", country.HardDeleteCountry).Methods("DELETE")

	// Swagger
	r.PathPrefix("/api/documentation/").Handler(httpSwagger.Handler(
		httpSwagger.URL(envVars.Scheme+"://"+envVars.Host + ":" + envVars.Port+"/api/documentation/doc.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	return r

}
