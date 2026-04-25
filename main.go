package main

import (
	//"database/sql"
	"fmt"
	"log"
	"svetlana/first-app/config"
	"svetlana/first-app/practice"

	// lets add .env
	//"os"

	"github.com/joho/godotenv"

	// This is the driver we just "got"
	//In Go, _ means “import for side effects only”
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	for name, driver := range config.ConnectionNames {
		db, err := config.InitConnection(name, driver)
		if err != nil {
			// we want to stop here
			// why?
			// if one or more db connections not ready app should not bootstrap
			log.Fatalf("Connection %s not ready! Error: %#v", name, err)
		}
		fmt.Printf("Connection '%s' is ready!\n", name)
		//config.PingConn(name)
		fmt.Printf("\u2705 Connection '%s' was 📌 added to map and is ready!\n", name)
		defer db.Close()
	}
	practice.RunLearning()
}
