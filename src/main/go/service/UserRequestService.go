package service

import (
	"context"
	"go-reviewer-api-service/src/main/go/model"
	"time"
)

type UserRequestService struct {
	userRequestRepo model.UserRequestRepository
	contextTimeout  time.Duration
}

func (u UserRequestService) GetAll(c context.Context) (res []model.UserRequest, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRequestRepo.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

func (u UserRequestService) GetUserRequestByOrgId(c context.Context, id int64) (res model.UserRequest, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err = u.userRequestRepo.GetUserRequestByOrgId(ctx, id)
	if err != nil {
		return
	}
	return
}

func NewUserRequestService(userRequestRepo model.UserRequestRepository, timeout time.Duration) model.UserRequestService {
	return &UserRequestService{
		userRequestRepo: userRequestRepo,
		contextTimeout:  timeout,
	}
}
