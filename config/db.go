package config

import (
	"github.com/freitagsrunde/protokollamt/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// OpenDatabase attempts to connect to configured
// PostgreSQL database and checks for connectivity.
func OpenDatabase(connString string, deployStage string, resetProt bool, resetRemv bool, resetRepl bool) (*gorm.DB, error) {

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

	// If it was indicated to drop the protocols table,
	// execute corresponding database command.
	if resetProt {
		db.DropTableIfExists(&models.Protocol{})
	}

	// If it was indicated to drop the removals table,
	// execute corresponding database command.
	if resetRemv {
		db.DropTableIfExists(&models.Removal{})
	}

	// If it was indicated to drop the replacements table,
	// execute corresponding database command.
	if resetRepl {
		db.DropTableIfExists(&models.Replacement{})
	}

	// Check if required tables are found.
	foundProt := db.HasTable(&models.Protocol{})
	foundRem := db.HasTable(&models.Removal{})
	foundRep := db.HasTable(&models.Replacement{})

	// If not, create them.
	if foundProt != true {
		db.CreateTable(&models.Protocol{})
	}

	if foundRem != true {
		db.CreateTable(&models.Removal{})
	}

	if foundRep != true {
		db.CreateTable(&models.Replacement{})
	}

	return db, nil
}
