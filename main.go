package main

import (
	"log"

	"github.com/ademspr/go-class/api"
)

func main() {

	log.Println("Teste do log")
	a := api.Server{}
	a.StartApi()
}
