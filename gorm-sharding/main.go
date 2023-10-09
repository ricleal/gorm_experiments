package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

type Order struct {
	ID        int64 `gorm:"primaryKey"`
	UserID    int64
	ProductID int64
	Name      string
}

const records = 128

func main() {
	log.Println("Starting")
	// Connect to database
	dbDSN := os.Getenv("DB_GORM_DSN")
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}

	for i := 0; i < 6; i += 1 {
		table := fmt.Sprintf("orders_%02d", i)
		db.Exec(`DROP TABLE IF EXISTS ` + table)
		db.Exec(`CREATE TABLE ` + table + ` (
			id bigint(20) unsigned NOT NULL,
			user_id bigint,
			product_id bigint,
			name varchar(255),
			PRIMARY KEY (id)
		)`)
	}

	middleware := sharding.Register(sharding.Config{
		ShardingKey:         "user_id",
		NumberOfShards:      6,
		PrimaryKeyGenerator: sharding.PKMySQLSequence,
		ShardingAlgorithm: func(columnValue any) (suffix string, err error) {
			if user_id, ok := columnValue.(int64); ok {
				return fmt.Sprintf("_%02d", user_id%6), nil
			}
			return "", errors.New("invalid user_id")
		},
	}, "orders")
	db.Use(middleware)

	// db.AutoMigrate(&Order{})

	for i := 0; i < records; i += 1 {

		err = db.Create(&Order{UserID: int64(i), ProductID: int64(i % 6),
			Name: fmt.Sprintf("name_%d", i)}).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	for i := 0; i < records; i += 1 {
		var orders []Order
		err = db.Model(&Order{}).Where("user_id", int64(i)).Find(&orders).Error
		if err != nil {
			fmt.Println("Error! Trying raw query")
			db.Raw("SELECT * FROM orders WHERE user_id = ?", int64(i)).Scan(&orders)
		}
		fmt.Printf("Orders: %+v\n", orders)
	}

}
