package config

import (
	"github.com/freitagsrunde/protokollamt/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// OpenDatabase attempts to connect to configured
// PostgreSQL database and checks for connectivity.
func OpenDatabase(connString string, deployStage string) (*gorm.DB, error) {

	var db *gorm.DB
	var err error

	// Open up a connection to configured database.
	db, err = gorm.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check connection to database in order to be sure.
	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}

	// If app runs in development mode, log SQL queries.
	if deployStage == "dev" {
		db.LogMode(true)
	}

	// Check if required tables are found.
	found := db.HasTable(&models.Protocol{})

	// If not, create them.
	if found != true {
		db.CreateTable(&models.Protocol{})
	}

	return db, nil
}
