package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"

	"os"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	addr := flag.String("addr", ":"+port, "TCP address to listen to")

	r := router.New()
	router.Mount(r)

	// Check mongoDB database connection
	//go model.CheckDatabaseConnection()

	// start the server
	fmt.Printf("API Service is running at port %v\n", *addr)
	if err := fasthttp.ListenAndServe(*addr, r.Handler); err != nil {
		log.Fatal(err)
	}

}
