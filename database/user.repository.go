package database

import (
	dbmodels "bluebird/models/dbmodels"
)

func GetUserByName(username string) (dbmodels.User, error) {
	db := GetDbConnect()
	db.Debug().LogMode(true)

	var user dbmodels.User
	var err error

	db.Where("username = ?", username).Find(&user)
	return user, err
}

func GetByUsername(username string) (dbmodels.User, error) {
	db := GetDbConnect()
	db.Debug().LogMode(true)

	var user dbmodels.User
	var err error

	db.Where("username = ?", username).First(&user)
	return user, err
}