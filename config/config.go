package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

const projectDirName = "gin-example" // change to relevant project name

func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
