package service

import (
	"context"
	"errors"
	"postman/internal/model"
	"postman/internal/storage"
	"postman/internal/storage/entity"
	"strconv"
)

type PostmanService struct {
	dbStorage *storage.DBStorage
}

func NewPostmanService(dbStorage *storage.DBStorage) *PostmanService {
	return &PostmanService{
		dbStorage: dbStorage,
	}
}

func (p *PostmanService) SimpleDBCreate(ctx context.Context, serviceRequest *model.ServiceRequest) (*model.ServiceResponse, error) {
	comment := serviceRequest.NewCommentEntityFromService()

	err := p.dbStorage.SaveComment(comment)

	if err != nil {
		return nil, err
	}

	return &model.ServiceResponse{CommentID: comment.ID}, nil
}

func (p *PostmanService) SimpleDBSearch(ctx context.Context, serviceRequest *model.ServiceRequest) (*model.ServiceResponse, error) {
	pageNo, err := strconv.Atoi(serviceRequest.Page)
	if err != nil {
		return nil, errors.New("invalid number")
	}

	comment := &entity.Comment{PostID: serviceRequest.PostID}

	comments, err := p.dbStorage.GetCommentPage(comment, pageNo)
	if err != nil {
		return nil, err
	}

	return &model.ServiceResponse{PostID: comment.PostID, Comments: comments}, nil
}
