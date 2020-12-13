package main

import (
	"net/http"
)

func main() {
	var s server
	s.ConfigFile = "config.json"
	s.Config = s.LoadConfig()
	s.Router = s.InitServer()
	err = http.ListenAndServe(":8084", s.Router)
}
