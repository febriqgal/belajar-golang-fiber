package migration

import (
	"belajar-golang-fiber/database"
	"belajar-golang-fiber/model/entity"
	"fmt"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		log.Println("Migration Failed")
	}
	fmt.Println("Migration Success")

}
