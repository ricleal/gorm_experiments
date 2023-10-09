package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	mysqlmigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
	AddressID int
	Address   Address
}

type Company struct {
	gorm.Model
	Name string
}

type Address struct {
	gorm.Model
	Street string
	City   string
	State  string
	Zip    string
}

func Migrate(gormDB *gorm.DB) error {
	db, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("error getting database from gorm: %w", err)
	}
	driver, err := mysqlmigrate.WithInstance(db, &mysqlmigrate.Config{})
	if err != nil {
		return fmt.Errorf("error creating migration driver: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql", driver)
	if err != nil {
		return fmt.Errorf("error creating migration instance: %w", err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to run")
			return nil
		}
		return fmt.Errorf("error running migrations: %w", err)
	}
	return nil
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
	// db.AutoMigrate(&User{}, &Company{})
	err = Migrate(db)
	if err != nil {
		log.Fatalln("Error migrating database:", err)
	}

	// Create Company
	log.Println("Creating Company")
	company := Company{Name: "Company"}
	db.Create(&company)

	// Create Address
	log.Println("Creating Address")
	address := Address{Street: "123 Main St", City: "Anytown", State: "TX", Zip: "12345"}
	db.Create(&address)

	// Create User
	log.Println("Creating User")
	user := User{Name: "User", CompanyID: int(company.ID), AddressID: int(address.ID)}
	db.Create(&user)

	log.Println("Querying User:")
	var user2 User
	db.Preload("Company").First(&user2, user.ID)
	log.Printf("%+v\n", user2)

	log.Println("Querying Company")
	var company2 Company
	db.First(&company2, company.ID)
	log.Printf("%+v\n", company2)

	log.Println("Querying Address")
	var address2 Address
	db.First(&address2, address.ID)
	log.Printf("%+v\n", address2)

	// Create User and Company in one transaction
	{
		log.Println("Creating User and Company in one transaction")
		tx := db.Begin()
		if tx.Error != nil {
			log.Fatalln("Error starting transaction:", tx.Error)
		}
		company3 := Company{Name: "Company3"}
		t := tx.Create(&company3)
		if t.Error != nil {
			tx.Rollback()
			log.Fatalln("Error creating company:", t.Error)
		}
		address3 := Address{Street: "456 Main St", City: "ZooTown", State: "MD", Zip: "45678"}
		t = tx.Create(&address3)
		if t.Error != nil {
			tx.Rollback()
			log.Fatalln("Error creating address:", t.Error)
		}
		user3 := User{Name: "User3", CompanyID: int(company3.ID), AddressID: int(address3.ID)}
		t = tx.Create(&user3)
		if t.Error != nil {
			tx.Rollback()
			log.Fatalln("Error creating user:", t.Error)
		}
		tx.Commit()
	}

	log.Println("Deleting User")
	db.Delete(&user)

	log.Println("Deleting Company")
	db.Delete(&company)

	log.Println("Deleting Address")
	db.Delete(&address)

	log.Println("Done")
}
