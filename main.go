package main

import (
	"log"

	"github.com/DarioKnezovic/campaign-service/routers"
)

func main() {
	r := routers.SetupRouter()


	log.Println("Server is running!")
	r.Run(":4000")
}