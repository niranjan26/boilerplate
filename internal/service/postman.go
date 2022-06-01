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
	user := &entity.User{
		Name: serviceRequest.UserName,
		Age:  serviceRequest.Age,
	}

	err := p.dbStorage.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return &model.ServiceResponse{ID: user.ID}, nil
}

func (p *PostmanService) SimpleDBSearch(ctx context.Context, serviceRequest *model.ServiceRequest) (*model.ServiceResponse, error) {
	id, err := strconv.Atoi(serviceRequest.ID)
	if err != nil {
		return nil, errors.New("invalid number")
	}

	user := &entity.User{
		ID: uint(id),
	}

	user, err = p.dbStorage.FindUserByID(user)
	if err != nil {
		return nil, err
	}

	return &model.ServiceResponse{ID: user.ID, Name: user.Name, Age: user.Age}, nil
}
