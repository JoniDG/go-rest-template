package main

import (
	"log"

	"github.com/JoniDG/go-rest-template/internal/router"
)

func main() {
	r := router.New()

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
