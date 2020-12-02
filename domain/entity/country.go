package entity

// Country represents the model for a country
type Country struct {
	Base
	Name      string	`json:"name"`	
	Iso       string	`json:"iso"`
	Iso3      string	`json:"iso3"`
	Numcode   string	`json:"numcode"`
	Phonecode string	`json:"phonecode"`
	Currency  string	`json:"currency"`
	Status    string	`json:"status"`
}