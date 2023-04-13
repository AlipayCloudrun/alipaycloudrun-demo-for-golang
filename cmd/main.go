package main

import (
	"log"
)

func main() {
	log.Println("starting alipaycloudrun-demo-for-go")
	server := InitServer()
	server.Start()
}
