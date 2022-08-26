package model

import (
	"context"
)

type UserRequest struct {
	Id             int `json:"id"`
	UserId         int `json:"user_id"`
	OrganizationId int `json:"organization_id"`
	Status         int `json:"status"`
}
type UserRequestService interface {
	GetAll(c context.Context) ([]UserRequest, error)
	GetUserRequestByOrgId(c context.Context, id int64) (UserRequest, error)
}

type UserRequestRepository interface {
	GetAll(c context.Context) ([]UserRequest, error)
	GetUserRequestByOrgId(c context.Context, id int64) (UserRequest, error)
}
