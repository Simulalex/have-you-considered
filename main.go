package main

import (
	"fmt"
	"github.com/Simulalex/haveyouconsidered/controllers"
	"log"
)

func main() {
	fmt.Println("Starting server...")

	config := controllers.Config{}
	//TODO: replace with environment variables or command line arguments or something
	config.Db.Username = "some-user"
	config.Db.Password = "some-password"
	config.Db.Database = "/haveyouconsidered"


	if err := controllers.Run(&config); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}
