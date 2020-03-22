package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func load_env() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env")
	}
}

func main() {
	load_env()
	fmt.Println(os.Getenv("API_KEY"))
}
