package main

import (
	"github.com/arxdsilva/knab/internal/middlewares"
	"github.com/arxdsilva/knab/platform/config"
	"github.com/prest/cmd"
	"github.com/prest/config/router"
	pms "github.com/prest/middlewares"
)

func main() {
	config.Load()
	pms.GetApp()
	r := router.Get()
	middlewares.RouterRegister(r)
	cmd.Execute()
}
