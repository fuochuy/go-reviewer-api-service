package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	error2 "go-reviewer-api-service/src/main/go/error"
	"go-reviewer-api-service/src/main/go/model"
)

type MySqlUserRequestRepository struct {
	Conn *sql.DB
}

func NewMySqlUserRequestRepository(Conn *sql.DB) model.UserRequestRepository {
	return &MySqlUserRequestRepository{Conn}
}

func (m *MySqlUserRequestRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.UserRequest, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]model.UserRequest, 0)
	for rows.Next() {
		t := model.UserRequest{}
		organizationId := int64(0)
		err = rows.Scan(
			&t.Id,
			&t.UserId,
			&organizationId,
			&t.Status,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.OrganizationId = int(organizationId)
		result = append(result, t)
	}

	return result, nil
}

func (m *MySqlUserRequestRepository) GetUserRequestByOrgId(c context.Context, id int64) (res model.UserRequest, err error) {
	query := `SELECT * FROM user_requests WHERE ID = ?`

	list, err := m.fetch(c, query, id)
	if err != nil {
		return res, err
	}

	if len(list) > 0 {
		return list[0], err
	} else {
		return res, error2.ErrNotFound
	}

}

func (m *MySqlUserRequestRepository) GetAll(c context.Context) (res []model.UserRequest, err error) {
	query := `SELECT * FROM user_requests`
	list, err := m.fetch(c, query)

	if err != nil {
		return res, err
	}
	if len(list) > 0 {
		return list, err
	} else {
		return list, error2.ErrNotFound
	}
}
