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
	resetProtFlag := flag.Bool("resetProtocols", false, "Append this flag in order to drop protocols table to start fresh.")
	resetRemvFlag := flag.Bool("resetRemovals", false, "Append this flag in order to drop removals table to start fresh.")
	resetReplFlag := flag.Bool("resetReplacements", false, "Append this flag in order to drop replacements table to start fresh.")
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
		log.Fatalf("Failed to load environment file: %v", err)
	}

	// Load protokollamt configuration file.
	c, err := config.LoadConfig(configName, os.Getenv("DB_PASSWORD"), os.Getenv("MAIL_PASSWORD"))
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to configured database and make sure
	// required databases exist.
	c.Database.Conn, err = config.OpenDatabase(c.Database.ConnString, c.DeployStage, *resetProtFlag, *resetRemvFlag, *resetReplFlag)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Initialize routes.
	router := config.DefineRoutes(c)

	// Run application.
	log.Printf("Protokollamt awaiting requests on %s.\n", c.PublicAddr)
	err = router.Run(c.PublicAddr)
	if err != nil {
		log.Fatalf("Running protokollamt failed: %v", err)
	}
}
