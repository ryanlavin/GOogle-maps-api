package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"github.com/gorilla/mux"
)

var err error

type server struct {
	Config *Conf
	ConfigFile string
	Router *mux.Router
}

type Conf struct {
	IndexURL string `json: "IndexURL"`
	IndexFile string `json: "IndexFile"`
	MapsURL string `json: "MapsURL"`
	MapsFile string `json: "MapsFile"`
}

func (s server) InitServer() *mux.Router {
	s.Router = mux.NewRouter()
	s.Router.HandleFunc(s.Config.IndexURL, s.Index)
	s.Router.HandleFunc(s.Config.MapsURL, s.Maps)
	return s.Router
}

func (s server) LoadConfig() *Conf {
	jsonFile, err := os.Open(s.ConfigFile)
	if err != nil {
		panic(err.Error)	
	}
	defer jsonFile.Close()

	confData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err.Error)
	}
	json.Unmarshal(confData, &s.Config)
	return s.Config
}

func (s server) Index(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, s.Config.IndexFile)	
}

func ShowImage(res http.ResponseWriter, req *http.Request) {
	address := req.FormValue("str")
	safeAddr := url.QueryEscape(address)
	fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?address=%s", safeAddr)

	c := appengine.NewContext(req)
	client := urlfetch.Client(c)
	resp, err := client.Get(fullUrl)
	if err != nil {
		panic(err.Error())
	}
        
        resp, err := client.Get(fullUrl)
}

func (s server) Maps(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		http.ServeFile(res, req, s.Config.MapsFile)
		return
	}
}
