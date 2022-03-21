package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

// env ...
var env ENV

// projectDirName ...
const projectDirName = "MongodbTest"

// InitDotEnv ...
func InitDotEnv() {
	// get env path
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))
	envPath := string(rootPath) + `/.env`

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("Error loading .env file")
	}

	// database ...
	database := Database{Uri: GetEnvString("DB_URI"), Name: GetEnvString("DB_Name"), TestName: GetEnvString("DB_Name_Test")}

	// appPort ...
	appPort := GetEnvString("APP_PORT")

	env = ENV{
		Database: database,
		AppPort:  appPort,
	}

}

// GetEnvString ...
func GetEnvString(key string) string {
	return os.Getenv(key)
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
