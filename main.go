package main

import (
	"github.com/ademspr/go-class/api"
	"github.com/ademspr/go-class/config"
)

func main() {
	c := config.Configuration{}
	c.Load()

	// l := logger.Logger{}
	// l.Init(c.Log)
	// defer l.Finish()

	a := api.Server{}
	a.Start(c.Server)
}
