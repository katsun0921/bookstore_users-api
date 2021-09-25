package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/katsun0921/bookstore_users-api/utils/errors"
	"strings"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database mysql response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data error")
	}
	return errors.NewInternalServerError("error processing request")
}
