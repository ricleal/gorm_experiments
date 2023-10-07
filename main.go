package main

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	gorm.Model
	Name string
}

func main() {

	log.Println("Starting")
	// Connect to database
	dbDSN := os.Getenv("DB_GORM_DSN")
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}

	// Migrate the schema
	db.AutoMigrate(&User{}, &Company{})

	// Create Company
	log.Println("Creating Company")
	company := Company{Name: "Company"}
	db.Create(&company)

	// Create User
	log.Println("Creating User")
	user := User{Name: "User", CompanyID: int(company.ID)}
	db.Create(&user)

	log.Println("Querying User:")
	var user2 User
	db.Preload("Company").First(&user2, user.ID)
	log.Printf("%+v\n", user2)

	log.Println("Querying Company")
	var company2 Company
	db.First(&company2, company.ID)
	log.Printf("%+v\n", company2)

	log.Println("Deleting User")
	db.Delete(&user)

	log.Println("Deleting Company")
	db.Delete(&company)

	log.Println("Done")
}
