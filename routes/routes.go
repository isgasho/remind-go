package routes

import (
	"net/http"
	"remind-go/handlers"
)

type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type WebRoutes []WebRoute

var webRoutes = WebRoutes{
	{
		"home",
		"GET",
		"/",
		handlers.Message,
	},
	{
		"home",
		"GET",
		"/home",
		handlers.Hello,
	},
}
