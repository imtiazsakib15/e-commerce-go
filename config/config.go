package config

import (
	"log"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	HTTPPort    int
	JwtSecret   string
}

var config Config

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if(version == "") {
		log.Println("Version is required!")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if(serviceName == "") {
		log.Println("Service name is required!")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if(httpPort == "") {
		log.Println("HTTP Port is required!")
		os.Exit(1)
	}

	JwtSecret := os.Getenv("JWT_SECRET")
	if(JwtSecret == "") {
		log.Println("Jwt secret is required!")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if(err != nil) {
		log.Println("HTTP Port is invalid!")
		os.Exit(1)
	}

	config = Config{
		Version:     version,
		ServiceName: serviceName,
		HTTPPort:    port,
	}
}

func GetConfig() Config {
	loadConfig()
	return config
}