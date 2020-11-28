package main

import (
	"github.com/arxdsilva/knab/internal/middlewares"
	"github.com/prest/cmd"
	"github.com/prest/config"
	"github.com/prest/config/router"
)

func main() {
	config.Load()
	r := router.Get()
	middlewares.RouterRegister(r)
	cmd.Execute()
}
