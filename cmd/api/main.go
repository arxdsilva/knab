package main

import (
	"github.com/arxdsilva/knab/internal/middlewares"
	"github.com/arxdsilva/knab/platform/config"
	"github.com/prest/prest/cmd"
	"github.com/prest/prest/config/router"
	pms "github.com/prest/prest/middlewares"
)

func main() {
	config.Load()
	pms.GetApp()
	r := router.Get()
	middlewares.Load()
	middlewares.RouterRegister(r)
	cmd.Execute()
}
