package repository

import (
	"fmt"
	"github.com/alphablaster/mocker/pkg/entity"
	"github.com/jmoiron/sqlx"
	"strings"
)

type MockPostgres struct {
	db *sqlx.DB
}

func NewMockPostgres(db *sqlx.DB) *MockPostgres {
	return &MockPostgres{db: db}
}

func (r *MockPostgres) Create(mock entity.Mock) (int, error) {
	var id int
	mockFields := []string {
		"name",
		"description",
		"status",
		"content_type",
		"request_method",
		"route_path",
		"body_type",
		"body_content",
		"script_type",
		"script",
		"active",
		"mock_order",
	}
	mockFieldsString := strings.Join(mockFields[:], ", ")

	query := fmt.Sprintf("INSERT INTO %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id",
		mockFieldsString, mocksTable)

	row := r.db.QueryRow(query, mock.Name, mock.Description, mock.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
