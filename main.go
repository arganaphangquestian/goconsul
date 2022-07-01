package main

import (
	"goconsul/utils"
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := api.NewClient(&api.Config{
		Address: "127.0.0.1:8500", // <-- Change this Address to Consul Server
		Scheme:  "http",
	})
	if err != nil {
		panic(err)
	}
	f := utils.FeatureFlag{Client: client} // Initiate Feature Flag
	if f.Get("arithmetic/features/addition") {
		log.Println("Addition", addition(2, 2))
	}
	if f.Get("arithmetic/features/substract") {
		log.Println("Substract", addition(2, 7))
	}
	if f.Get("arithmetic/features/divide") {
		log.Println("Divide", addition(9, 2))
	}
	if f.Get("arithmetic/features/multiply") {
		log.Println("Multiply", addition(3, 2))
	}
}

func addition(a, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func divide(a, b float32) float32 {
	return a / b
}

func multiply(a, b int) int {
	return a * b
}
