package main

import (
	"flag"
	"log"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
	"gitlab.com/visaruttirataworawan/grader_gw/router"
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

	log.Println("API service listened on port ",*addr)
	if err := fasthttp.ListenAndServe(*addr, r.Handler); err != nil {
		log.Fatal(err)
	}

}
