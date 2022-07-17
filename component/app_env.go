package component

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() {
	env := os.Getenv("ENV")
	envFile := ".env"

	if env == "DOCKER" {
		envFile = ".docker.env"
	}

	log.Printf("env: %s - envFile : %s \n", env, envFile)

	if err := godotenv.Load(envFile); err != nil {
		log.Fatalln(err)
		//panic(err)
	}

	log.Println("LoadEnv completed ...")
}
