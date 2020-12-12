package models

import (
	"github.com/jinzhu/gorm"
	"io"
)

type ClientApi struct {
	gorm.Model

	Environment  string    `db:"environment"`
	ClientName   string    `db:"client_name"`
	Url          string    `db:"url"`
	RequestBody  io.Reader `db:"request_body"`
	ResponseBody string    `db:"response_body"`
	HttpCode     int       `db:"status"`
}
