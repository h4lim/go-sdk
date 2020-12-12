package models

import (
	"github.com/jinzhu/gorm"
)

type LogApi struct {
	gorm.Model

	Environment  string `db:"environment"`
	ClientName   string `db:"client_name"`
	Url          string `db:"url"`
	Method       string `db:"method"`
	Header       string `db:"header"`
	RequestBody  string `db:"request_body"`
	ResponseBody string `db:"response_body"`
	HttpCode     int    `db:"status"`
}
