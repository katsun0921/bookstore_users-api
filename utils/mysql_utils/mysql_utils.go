package mysql_utils

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/katsun0921/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_errors.NewNotFoundError("no record matching given id")
		}
		return rest_errors.NewInternalServerError("error parsing database mysql response", err)
	}

	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewBadRequestError("invalid data error")
	}
	return rest_errors.NewInternalServerError("error processing request", errors.New("database error"))
}
