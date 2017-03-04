package main

import (
	"flag"
	"log"
	"os"

	"github.com/freitagsrunde/protokollamt/config"
	"github.com/joho/godotenv"
)

func main() {

	// Allow to overwrite default names of
	// environment and configuration files
	// for protokollamt.
	envNameFlag := flag.String("envName", ".env", "Define name of protokollamt environment file. Providing a path will fail!")
	configNameFlag := flag.String("configName", "config.toml", "Define name of protokollamt configuration file. Providing a path will fail!")
	flag.Parse()

	// Check that flag values are no paths.
	envName, configName, err := config.CheckFlags(*envNameFlag, *configNameFlag)
	if err != nil {
		log.Fatalf("Error while checking flags: %v", err)
	}

	// Load .env file for deployment-specific
	// environment values.
	err = godotenv.Load(envName)
	if err != nil {
		log.Fatal("Failed to load environment file: %v", err)
	}

	// Load protokollamt configuration file.
	_, err = config.LoadConfig(configName, os.Getenv("DB_PASSWORD"), os.Getenv("MAIL_PASSWORD"))
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Check setup: LDAP connection, database
	// connection, and mail sending capabilities.

	// Initialize routes.

	// Run application.
}
