package main

import (
	"fmt"
	"github.com/dhowden/numerus"
	"log"
)

func main() {
	num, err := numerus.Parse("V")
	if err != nil {
		log.Println("error parsing", err)
	}
	fmt.Println("parse roman num", num)
}
