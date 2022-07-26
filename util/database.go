package util

import (
	"fmt"
	"log"

	"github.com/RaguSoftEng/inviteservice_api/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db_env := AppConfig.Db

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_env.Username, db_env.Password, db_env.Host, db_env.Port, db_env.Databse)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("[ ERROR ] Unable to connect with mysql!\n", err)
	}

	fmt.Println("[ OK ] Connected to the DB!")

	return db

}

func CloseDbConnection(con *gorm.DB) {
	sql, err := con.DB()

	if err != nil {
		log.Fatal("[ ERROR ] Unable to close the database connection.")
	}

	sql.Close()
}

func Migrate() {
	db := Connect()

	if err := db.AutoMigrate(&models.User{}, &models.InviteToken{}); err != nil {
		log.Fatal("[ ERROR ] Unable to migrate!\n", err)
		return
	}
	log.Println("[ OK ] Database migrated successfully.")
}
