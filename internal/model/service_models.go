package model

import "postman/internal/storage/entity"

type ServiceRequest struct {
	Page      string
	CommentID uint
	PostID    uint
	UserID    uint
	ParentID  uint
	Comment   string
	UserName  string
	Age       float64
	ID        string
}

type ServiceResponse struct {
	ID        uint              `json:"ID,omitempty"`
	Name      string            `json:"name,omitempty"`
	Age       float64           `json:"age,omitempty"`
	CommentID uint              `json:"commentID,omitempty"`
	PostID    uint              `json:"postID,omitempty"`
	UserID    uint              `json:"userID,omitempty"`
	ParentID  uint              `json:"parentID,omitempty"`
	Comment   string            `json:"comment,omitempty"`
	Comments  []*entity.Comment `json:"comments,omitempty"`
}

func (s *ServiceRequest) NewCommentEntityFromService() *entity.Comment {
	return &entity.Comment{
		PostID:   s.PostID,
		UserID:   s.UserID,
		ParentID: s.ParentID,
		Comment:  s.Comment,
	}
}
