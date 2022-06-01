package service

import (
	"context"
	"postman/internal/model"
)

type Service interface {
	SimpleDBCreate(ctx context.Context, serviceRequest *model.ServiceRequest) (*model.ServiceResponse, error)
	SimpleDBSearch(ctx context.Context, serviceRequest *model.ServiceRequest) (*model.ServiceResponse, error)
}
