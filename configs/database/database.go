package database

import (
	"log"
	// "os"
	// "strconv"
	// "time"

	"github.com/denyherianto/go-fiber-boilerplate/app/utils"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// _ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
)

var DBManager *gorm.DB

func OpenDBConnection() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// Define a new Database connection.
	// Define database connection settings.
	// maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))
	// maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	// maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// Build PostgreSQL connection URL.
	dsn, dsnErr := utils.ConnectionURLBuilder("postgres")
	if dsnErr != nil {
		log.Println(dsnErr)
	}

	// Define database connection for PostgreSQL.
	// https://github.com/go-gorm/postgres
	DBManager, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}
	log.Println("Database connected", DBManager.Name())

	// sqlDB, sqlDBErr := DBManager.DB()

	// if sqlDBErr != nil {
	// 	log.Println(sqlDBErr)
	// }

	// Set database connection settings:
	// 	- SetMaxOpenConns: the default is 0 (unlimited)
	// 	- SetMaxIdleConns: defaultMaxIdleConns = 2
	// 	- SetConnMaxLifetime: 0, connections are reused forever
	// sqlDB.SetMaxIdleConns(maxConn)
	// sqlDB.SetMaxOpenConns(maxIdleConn)
	// sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))
}
