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
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write([]byte("{\"result\":\"OK\"}"))
		},
	},
}
