package service

import "net/http"

// Route --
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes --
type Routes []Route

var routes = Routes{
	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountID}",
		GetAccount,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}
