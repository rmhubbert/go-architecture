package main

import (
	"log"

	"github.com/rmhubbert/rmhttp"
)

func main() {
	rmh := rmhttp.New()
	InitRoutes(rmh)

	log.Fatal(rmh.Start())
}
