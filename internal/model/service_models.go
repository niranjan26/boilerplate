package model

type ServiceRequest struct {
	UserName string
	Age      float64
	ID       string
}

type ServiceResponse struct {
	ID   uint    `json:"ID"`
	Name string  `json:"name"`
	Age  float64 `json:"age"`
}
