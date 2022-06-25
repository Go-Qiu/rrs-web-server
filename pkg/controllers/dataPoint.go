package controllers

type DataPoint struct {
	UserID    float64 `json:"id"`
	Phone     string  `json:"phone"`
	Name      string  `json:"name"`
	Points    float64 `json:"points"`
	LastLogin string  `json:"last_login"`
}

type DataPointExtended struct {
	UserID    float64 `json:"id"`
	Phone     string  `json:"phone"`
	Name      string  `json:"name"`
	Password  string  `json:"password"`
	Points    float64 `json:"points"`
	LastLogin string  `json:"last_login"`
}
