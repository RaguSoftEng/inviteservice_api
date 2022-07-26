package repository

import (
	"github.com/RaguSoftEng/inviteservice_api/src/models"
	"github.com/RaguSoftEng/inviteservice_api/util"
	"gorm.io/gorm"
)

//Seed users to database
// All users are considered as admin users
func ImportUsers(db *gorm.DB) {

	var user models.User
	result := db.First(&user)

	if result.RowsAffected > 0 {
		return
	}

	var users = []models.User{
		{FullName: "Nagarajah Raguvaran", Username: "Ragu", Password: "Ragu321"},
		{FullName: "Thankavel Ramesh", Username: "Ramesh", Password: "Ramesh001"},
		{FullName: "Ravivarman", Username: "Ravi", Password: "Ravi963"},
		{FullName: "William", Username: "william", Password: "123654"}}

	db.Create(&users)
}

// Validate user with username and password to login
func ValidateUser(db *gorm.DB, username string, password string) (string, error) {
	user := models.User{}
	var err error
	err = db.Model(models.User{}).Where("username = ? AND password = ?", username, password).Take(&user).Error

	if err != nil {
		return "", err
	}

	return util.CreateToken(user.ID)
}
