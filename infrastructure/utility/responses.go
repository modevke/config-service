package utility

type ResponseError struct{
	Code			string	`json:"code"`
	Description	string	`json:"description"`
}

type BaseResponse struct{
	Status		string		`json:"status"`
	Description	string		`json:"description"`
}

type SuccessResponse struct{
	BaseResponse
	Data		interface{}	`json:"data,omitempty"`
}

type ErrorResponse struct{
	BaseResponse
	Errors		[]ResponseError	`json:"error"`
}



type Responses struct{
	Status		string		`json:"status"`
	Description	string		`json:"description"`
	Data		interface{}	`json:"data,omitempty"`
	Errors		[]ResponseError	`json:"error"`
}

func (r *Responses) ErrorResponse(desc string, data []ResponseError ){
	r.Status = "001"
	r.Description = desc
	r.Errors = data
}

func (r *Responses) SuccessResponse(status string, desc string, data interface{} ){
	r.Status = status
	r.Description = desc
	r.Data = data
}