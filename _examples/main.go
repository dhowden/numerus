package main

import (
	"fmt"
	"log"
	"numerus"
)

func main() {
	num, err := numerus.Parse("VI")
	if err != nil {
		log.Println("error parsing", err)
	}
	fmt.Println("parse roman num", num)
}
