package controllers

// DataPoint holds the user data.
type DataPoint struct {
	UserID    float64 `json:"id"`
	Phone     string  `json:"phone"`
	Name      string  `json:"name"`
	Points    float64 `json:"points"`
	LastLogin string  `json:"last_login"`
}

// DataPointExtended holds the user data, with password hash included.
type DataPointExtended struct {
	UserID    float64 `json:"id"`
	Phone     string  `json:"phone"`
	Name      string  `json:"name"`
	Password  string  `json:"password"`
	Points    float64 `json:"points"`
	LastLogin string  `json:"last_login"`
}

// Transaction struct holds the transaction data.
type Transaction struct {
	Item   string  `json:"item"`
	Phone  string  `json:"phone"`
	Points float64 `json:"points"`
	Weight float64 `json:"weight"`
}

// ResponseBody struct holds the response data returned by microservice.
type ResponseBody struct {
	Ok   bool                   `json:"ok"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}
