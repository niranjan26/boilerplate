package model

import "errors"

type RequestDTO interface {
	Validate() error
	GetServiceRequest() *ServiceRequest
}

type PostmanRequest struct {
	UserName string  `json:"name"`
	Age      float64 `json:"age"`
}

func (p *PostmanRequest) Validate() error {
	if p.Age < 0.0 {
		return errors.New("invalid age")
	}

	return nil
}

func (p *PostmanRequest) GetServiceRequest() *ServiceRequest {
	return &ServiceRequest{
		UserName: p.UserName,
		Age:      p.Age,
	}
}

type Response struct {
	ResponseCode    string      `json:"ResponseCode"`
	ResponseMessage string      `json:"ResponseMessage"`
	ResponseVersion string      `json:"ResponseVersion"`
	ResponseData    interface{} `json:"Data"`
}
