package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	error2 "go-reviewer-api-service/src/main/go/error"
	"go-reviewer-api-service/src/main/go/model"
	"go-reviewer-api-service/src/main/go/response"
	"net/http"
	"strconv"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserRequestController struct {
	userRequestService model.UserRequestService
}

func NewUserRequestController(e *echo.Echo, us model.UserRequestService) {
	handler := &UserRequestController{
		userRequestService: us,
	}
	e.GET("/user-request/:id", handler.GetUserRequestByOrgId)
	e.GET("/user-requests", handler.GetAll)

}

func (u *UserRequestController) GetUserRequestByOrgId(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, error2.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	userRequet, err := u.userRequestService.GetUserRequestByOrgId(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	var res response.Response
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = userRequet
	return c.JSON(http.StatusOK, res)
}
func (u *UserRequestController) GetAll(c echo.Context) error {

	ctx := c.Request().Context()
	userRequets, err := u.userRequestService.GetAll(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}
	var res response.Response
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = userRequets
	return c.JSON(http.StatusOK, res)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case error2.ErrInternalServerError:
		return http.StatusInternalServerError
	case error2.ErrNotFound:
		return http.StatusNotFound
	case error2.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
