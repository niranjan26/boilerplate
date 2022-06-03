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

type PostmanCommentRequest struct {
	PostID   uint   `json:"postID"`
	UserID   uint   `json:"userID"`
	ParentID uint   `json:"parentID"`
	Comment  string `json:"comment"`
	Page     uint   `json:"page"`
}

func (p *PostmanCommentRequest) Validate() error {

	return nil
}

func (p *PostmanCommentRequest) GetServiceRequest() *ServiceRequest {
	return &ServiceRequest{
		Comment:  p.Comment,
		ParentID: p.ParentID,
		PostID:   p.PostID,
		UserID:   p.UserID,
	}
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
