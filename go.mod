module gorm-exp

go 1.21.0

require (
	github.com/golang-migrate/migrate/v4 v4.16.2
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.4
)

require gorm.io/sharding v0.0.0

replace gorm.io/sharding v0.0.0 => ../sharding-fork

require (
	github.com/bwmarrin/snowflake v0.3.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/longbridgeapp/sqlparser v0.3.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/exp v0.0.0-20230817173708-d852ddb80c63 // indirect
)
