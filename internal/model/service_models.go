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
	ID        uint              `json:"ID"`
	Name      string            `json:"name"`
	Age       float64           `json:"age"`
	CommentID uint              `json:"commentID"`
	PostID    uint              `json:"postID"`
	UserID    uint              `json:"userID"`
	ParentID  uint              `json:"parentID"`
	Comment   string            `json:"comment"`
	Comments  []*entity.Comment `json:"comments"`
}

func (s *ServiceRequest) NewCommentEntityFromService() *entity.Comment {
	return &entity.Comment{
		PostID:   s.PostID,
		UserID:   s.UserID,
		ParentID: s.ParentID,
		Comment:  s.Comment,
	}
}
