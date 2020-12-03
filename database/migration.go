package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/h4lim/go-sdk/logging"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	INTERNAL       = "INTERNAL"
	POSGRES_CONFIG = "user=%s password=%s dbname=%s host=%s port=%s sslmode=%s"
	MYSQL_CONFIG   = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
)

var log = logging.MustGetLogger("go-sdk")

type DBOperation interface {
	InitDB() (*gorm.DB, *error)
	DBOpen() (*gorm.DB, *error)
	Select(query string, args ...interface{}) (*gorm.DB, *sql.Rows, *error)
	Insert(query string, args ...interface{}) (*gorm.DB, *error)
	Update(query string, args ...interface{}) (*gorm.DB, *error)
	Delete(query string, args ...interface{}) (*gorm.DB, *error)
}

type DBModel struct {
	ServerMode string
	Driver     string
	Host       string
	Port       string
	Name       string
	Username   string
	Password   string
}

func (c *DBModel) InitDB() (*gorm.DB, *error) {
	db, err := dBOpen(c)
	if err != nil {
		log.Errorf(INTERNAL, "Error When Open DB %s ", err)
		return nil, err
	}
	return db, nil
}

func (c *DBModel) DBOpen() (*gorm.DB, *error) {
	db, err := dBOpen(c)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func dBOpen(c *DBModel) (*gorm.DB, *error) {

	var connectionUrl string
	switch c.Driver {
	case "postgres":
		connectionUrl = fmt.Sprintf(POSGRES_CONFIG, c.Username, c.Password, c.Name, c.Host, c.Port, "disable")
	case "mysql":
		connectionUrl = fmt.Sprintf(MYSQL_CONFIG, c.Username, c.Password, c.Host, c.Port, c.Name)
	default:
		log.Errorf(logging.INTERNAL, "No Database Selected!, Please check config.toml")
		os.Exit(1)
	}

	db, err := gorm.Open(c.Driver, connectionUrl)
	if err != nil {
		log.Errorf(logging.INTERNAL, "Cannot Connect to DB With Message ", err.Error())
		return nil, &err
	}

	return db, nil
}
