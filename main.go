package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "time/tzdata"
	"uploadCSV/server"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

func main() {
	srv := server.SrvInit()
	srv.Start(":8080")

	LoadEnv()
	fmt.Println("connected")
}
