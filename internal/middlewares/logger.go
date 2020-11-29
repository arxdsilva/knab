package middlewares

import (
	"github.com/prest/prest/middlewares"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

// Load loads middlewares to negroni's stack
func Load() {
	middlewares.MiddlewareStack = append(middlewares.MiddlewareStack, negroni.NewRecovery())
	middlewares.MiddlewareStack = append(middlewares.MiddlewareStack, negroni.NewLogger())
	middlewares.MiddlewareStack = append(middlewares.MiddlewareStack, middlewares.HandlerSet())
	middlewares.MiddlewareStack = append(middlewares.MiddlewareStack, cors.Default())
}
