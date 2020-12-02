package interfaces

import(
	// "log"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gorilla/mux"

	"github.com/modevke/config-service/domain/entity"
	"github.com/modevke/config-service/infrastructure/utility"
)

// SAMPLE DATA
var sample []entity.Country = []entity.Country{
	entity.Country{
		Base: entity.Base{
			ID: "123",
		},
		Name: "Kenya",
		Iso: "KE",
		Iso3: "KEN",
		Numcode: "404",
		Phonecode: "254",
		Currency: "KES",
	},
	entity.Country{
		Base: entity.Base{
			ID: "456",
		},
		Name: "Uganda",
		Iso: "UG",
		Iso3: "UGN",
		Numcode: "406",
		Phonecode: "256",
		Currency: "UGX",
	},
}

type CountryHandler struct{
	Environment	string
}

// CREATE COUNTRY
type CountryRequestSchema struct{
	Name      string	`json:"name"`	
	Iso       string	`json:"iso"`
	Iso3      string	`json:"iso3"`
	Numcode   string	`json:"numcode"`
	Phonecode string	`json:"phonecode"`
	Currency  string	`json:"currency"`
}
type CountryResposeSchema struct{
	utility.BaseResponse
	Data	entity.Country	`json:"data,omitempty"`
}
// CreateCountry godoc
// @Summary Create a new country
// @Tags country
// @Accept  json
// @Produce  json
// @Param  body body CountryRequestSchema false " "
// @Success 200 {object} CountryResposeSchema
// @Failure 400,404 {object} utility.ErrorResponse
// @Router /country/create [post]
func(c *CountryHandler) CreateCountry(res http.ResponseWriter, req *http.Request){

	var response utility.Responses
	var reqBody entity.Country

	am := utility.ValidateBody(req.Body, &reqBody, 
		validation.Field(&reqBody.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&reqBody.Iso, validation.Required, validation.Length(1, 2)),
		validation.Field(&reqBody.Iso3, validation.Required, validation.Length(1, 3)),
		validation.Field(&reqBody.Numcode, validation.Required, validation.Length(1, 100)),
		validation.Field(&reqBody.Phonecode, validation.Required, validation.Length(1, 100)),
		validation.Field(&reqBody.Currency, validation.Required, validation.Length(1, 3)),
	)

	if len(am) > 0 {
		
		response.ErrorResponse("Create a country: Invalid credentials", am)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		
		return
	}

	marand := rand.Intn(999 - 100) + 100

	reqBody.ID = strconv.Itoa(marand)

	sample = append(sample, reqBody)

	response.SuccessResponse("000", "Create a country", reqBody)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)

}


// FetchCountryByID godoc
// @Summary Fetch country by ID
// @Tags country
// @Accept  json
// @Produce  json
// @Param  id path string true "Country ID"
// @Success 200 {object} CountryResposeSchema
// @Failure 400,404,500 {object} utility.ErrorResponse
// @Router /country/fetch-id/{id} [get]
func(c *CountryHandler) FetchCountryByID(res http.ResponseWriter, req *http.Request){

	var response utility.Responses
	vars := mux.Vars(req)

	reqParams :=  struct{
		ID	string	`json:"id"`
	}{
		ID: vars["id"],
	}

	am := utility.ValidateParams(&reqParams, 
		validation.Field(&reqParams.ID, validation.Required, validation.Length(1, 50)),
	)

	if len(am) > 0 {
		response.ErrorResponse("Fetch Country By ID: Invalid credentials", am)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return
	}


	var data entity.Country
	var found bool
	for _, val := range sample{
		if val.Base.ID == vars["id"] {
			found = true
			data = val
			break
		}
	}

	if !found {

		response.ErrorResponse("Cannot find country", make([]utility.ResponseError, 0))
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return

	}
	
	response.SuccessResponse("000", "Fetch Country By ID", data)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)

}

// FetchCountryByISO godoc
// @Summary Fetch country by ISO
// @Tags country
// @Accept  json
// @Produce  json
// @Param  iso path string true "Country ISO"
// @Success 200 {object} CountryResposeSchema
// @Failure 400,404,500 {object} utility.ErrorResponse
// @Router /country/fetch-iso/{iso} [get]
func(c *CountryHandler) FetchCountryByISO(res http.ResponseWriter, req *http.Request){
	var response utility.Responses
	vars := mux.Vars(req)

	reqParams := struct{
		Iso       string	`json:"iso"`
	}{
		Iso: vars["iso"],
	}

	am := utility.ValidateParams(&reqParams,
		validation.Field(&reqParams.Iso, validation.Required, validation.Length(1, 2)),
	)

	if len(am) > 0 {
		response.ErrorResponse("Fetch Country By ISO: Invalid credentials", am)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return
	}

	var data entity.Country
	var found bool
	for _, val := range sample{
		if val.Iso == vars["iso"] {
			found = true
			data = val
			break
		}
	}

	if !found {

		response.ErrorResponse("Cannot find country", make([]utility.ResponseError, 0))
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return

	}

	response.SuccessResponse("000", "Fetch Country By ISO", data)
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)


}

type CountriesResposeSchema struct{
	utility.BaseResponse
	Data	[]entity.Country	`json:"data"`
}
// FetchCountries godoc
// @Summary Fetch all countries
// @Tags country
// @Accept  json
// @Produce  json
// @Success 200 {object} CountriesResposeSchema
// @Failure 400,404,500 {object} utility.ErrorResponse
// @Router /country/fetch-all [get]
func(c *CountryHandler) FetchCountries(res http.ResponseWriter, req *http.Request){

	var response utility.Responses

	response.SuccessResponse("000", "Fetch Countries", sample)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)
}


type UpdateCountryResposeSchema struct{
	ID    string	`json:"id"`
	CountryRequestSchema
}
// UpdateCountry godoc
// @Summary Update Country by ID
// @Tags country
// @Accept  json
// @Produce  json
// @Param  body body UpdateCountryResposeSchema false " "
// @Success 200 {object} CountryResposeSchema
// @Failure 400,404 {object} utility.ErrorResponse
// @Router /country/update [put]
func(c *CountryHandler) UpdateCountry(res http.ResponseWriter, req *http.Request){
	var response utility.Responses
	var reqBody entity.Country

	am := utility.ValidateBody(req.Body, &reqBody,
		validation.Field(&reqBody.ID, validation.Required, validation.Length(1, 50)),
		validation.Field(&reqBody.Name, validation.Length(1, 50)),
		validation.Field(&reqBody.Iso, validation.Length(1, 2)),
		validation.Field(&reqBody.Iso3, validation.Length(1, 3)),
		validation.Field(&reqBody.Numcode, validation.Length(1, 100)),
		validation.Field(&reqBody.Phonecode, validation.Length(1, 100)),
		validation.Field(&reqBody.Currency, validation.Length(1, 3)),
	)

	if len(am) > 0 {
		response.ErrorResponse("Update Country: Invalid credentials", am)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return
	}

	var found bool
	for i, val := range sample{
		if val.Base.ID == reqBody.ID {
			found = true
			if reqBody.Name != ""{
				sample[i].Name = reqBody.Name 
			}
			if reqBody.Iso != ""{
				sample[i].Iso = reqBody.Iso 
			}
			if reqBody.Iso3 != ""{
				sample[i].Iso3 = reqBody.Iso3 
			}
			if reqBody.Numcode != ""{
				sample[i].Numcode = reqBody.Numcode 
			}
			if reqBody.Phonecode != ""{
				sample[i].Phonecode = reqBody.Phonecode 
			}
			if reqBody.Currency != ""{
				sample[i].Currency = reqBody.Currency 
			}

			reqBody = sample[i]

			break
		}
	}

	if !found {

		response.ErrorResponse("Cannot find country", make([]utility.ResponseError, 0))
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return

	}

	response.SuccessResponse("000", "Update Country", reqBody)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)

}


// SoftDeleteCountry godoc
// @Summary Soft delete Country by ID
// @Tags country
// @Accept  json
// @Produce  json
// @Param  body body utility.IDVariable false " "
// @Success 200 {object} utility.BaseResponse
// @Failure 400,404 {object} utility.ErrorResponse
// @Router /country/delete-soft [delete]
func(c *CountryHandler) SoftDeleteCountry(res http.ResponseWriter, req *http.Request){

	var response utility.Responses
	reqBody := struct{
		ID	string	`json:"id"`
	}{}

	am := utility.ValidateBody(req.Body, &reqBody,
		validation.Field(&reqBody.ID, validation.Required, validation.Length(1, 50)),
	)

	if len(am) > 0 {
		response.ErrorResponse("Soft Delete Country: Invalid credentials", am)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return
	}

	var found bool
	for i, val := range sample{
		if val.Base.ID == reqBody.ID {
			found = true
		
			sample[i].Base.DeletedAt = time.Now()
			
			break
		}
	}

	if !found {

		response.ErrorResponse("Cannot find country", make([]utility.ResponseError, 0))
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return

	}


	response.SuccessResponse("000", "Soft Delete Country", reqBody)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)

}


// HardDeleteCountry godoc
// @Summary Hard delete Country by ID
// @Tags country
// @Accept  json
// @Produce  json
// @Param  body body utility.IDVariable false " "
// @Success 200 {object} utility.BaseResponse
// @Failure 400,404 {object} utility.ErrorResponse
// @Router /country/delete-hard [delete]
func(c *CountryHandler) HardDeleteCountry(res http.ResponseWriter, req *http.Request){

	var response utility.Responses
	reqBody := struct{
		ID	string	`json:"id"`
	}{}

	am := utility.ValidateBody(req.Body, &reqBody,
		validation.Field(&reqBody.ID, validation.Required, validation.Length(1, 50)),
	)

	if len(am) > 0 {
		response.ErrorResponse("Hard Delete Country: Invalid credentials", am)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return
	}

	var found bool
	for i, val := range sample{
		if val.Base.ID == reqBody.ID {
			found = true
			sample = append(sample[:i], sample[i+1:]...)
			
			break
		}
	}

	if !found {

		response.ErrorResponse("Cannot find country", make([]utility.ResponseError, 0))
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(response)
		return

	}


	response.SuccessResponse("000", "Hard Delete Country", reqBody)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)


}