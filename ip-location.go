package main

import "github.com/ipipdotnet/datx-go"
import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"flag"
)


var city *datx.City
var err error

func main(){
	var datafile string
	flag.StringVar(&datafile, "datafile", "17monipdb.datx", "datafile path")
	var port string
	flag.StringVar(&port, "port", "8080", "server port")
	flag.Parse()
	flag.Usage()
	city ,err = datx.NewCity(datafile)
	if err == nil {
		http.HandleFunc("/location", location)
		http.HandleFunc("/", handler)
		log.Println("server is start on "+ port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}


}

type IpLocation struct {
	Country string `json:"country"`
	Province string `json:"province"`
	City string `json:"city"`
}


func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.FormValue("ip")
	if ip == "" {
		return
	}
	log.Println(ip)
	location, _ := city.Find(ip)
	ipLocation := &IpLocation{location[0], location[1], location[2]}
	bytes, _ := json.Marshal(ipLocation)
	s := string(bytes)
	w.Header().Add("Content-Type","application/json; charset=utf-8")
	fmt.Fprintf(w, s )

}

func location(w http.ResponseWriter, r *http.Request) {
	ip := r.FormValue("ip")
	if ip == "" {
		return
	}
	log.Println(ip)
	location, _ := city.FindLocation(ip)
	w.Header().Add("Content-Type","application/json; charset=utf-8")
	fmt.Fprintf(w, string(location.ToJSON()))

}
