package models

import (
	"github.com/jinzhu/gorm"
)

type LogApi struct {
	gorm.Model

	LogID        string `db:"log_id"`
	Environment  string `db:"environment"`
	ClientName   string `db:"client_name"`
	Url          string `gorm:"size:5000000"`
	Method       string `db:"method"`
	Header       string `gorm:"size:5000000"`
	RequestBody  string `gorm:"size:5000000"`
	ResponseBody string `gorm:"size:5000000"`
	HttpCode     int    `db:"status"`
}
