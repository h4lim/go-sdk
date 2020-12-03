package database

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

func (c *DBModel) Select(query string, args ...interface{}) (*gorm.DB, *sql.Rows, *error) {

	DB, err := dBOpen(c)
	if err != nil {
		return DB, nil, err
	}

	res, err2 := DB.Raw(query, args...).Rows()
	return DB, res, &err2
}

func (c *DBModel) Insert(query string, args ...interface{}) (*gorm.DB, *error) {

	DB, err := dBOpen(c)
	if err != nil {
		return DB, err
	}

	tx := DB.Begin()
	tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return DB, err
	}

	tx.Commit()
	return DB, nil
}

func (c *DBModel) Update(query string, args ...interface{}) (*gorm.DB, *error) {

	DB, err := dBOpen(c)
	if err != nil {
		return DB, err
	}

	tx := DB.Begin()
	tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return DB, err
	}

	tx.Commit()
	return DB, nil
}

func (c *DBModel) Delete(query string, args ...interface{}) (*gorm.DB, *error) {

	DB, err := dBOpen(c)
	if err != nil {
		return DB, err
	}

	DB.Exec(query, args...)
	if err != nil {
		return DB, err
	}

	return DB, nil
}
