package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/gorilla/mux"
	"gobot.io/x/gobot/platforms/raspi"
)

// Options is program options
type Options struct {
	Port         int
	SwitchConfig map[string]SwitchConfiguration `yaml:"switches"`
}

// SwitchConfiguration defines a turnout switch pin configuration
type SwitchConfiguration struct {
	OnOffPin     int `yaml:"on_off_pin"`
	DirectionPin int `yaml:"direction_pin"`
}

// TrainServer is the container for the HTTP server
type TrainServer struct {
	switches map[string]*TrainSwitch
}

// ToggleSwitch toggles switch :id on or off
func (ts *TrainServer) ToggleSwitch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sw, ok := ts.switches[vars["id"]]
	if ok {
		status, err := sw.Toggle()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error toggling: %s", err.Error())
			return
		}
		if status {
			fmt.Fprint(w, "1")
		} else {
			fmt.Fprint(w, "0")
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// TurnOn turns the switch with :id on
func (ts *TrainServer) TurnOn(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sw, ok := ts.switches[vars["id"]]
	if ok {
		if err := sw.On(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error turning switch on: %s", err.Error())
			return
		}
		fmt.Fprint(w, "1")
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// TurnOff turns the switch with :id off
func (ts *TrainServer) TurnOff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sw, ok := ts.switches[vars["id"]]
	if ok {
		if err := sw.Off(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error turning switch off: %s", err.Error())
			return
		}
		fmt.Fprint(w, "0")
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// SwitchStatus returns the given id's status
func (ts *TrainServer) SwitchStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sw, ok := ts.switches[vars["id"]]
	if ok {
		if sw.Status() {
			fmt.Fprint(w, "1")
		} else {
			fmt.Fprint(w, "0")
		}
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	opts := Options{
		Port: 8080,
	}

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s /path/to/config", os.Args[0])
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("error reading config file: %s", err.Error())
	}

	if err := yaml.Unmarshal(b, &opts); err != nil {
		log.Fatalf("error parsing config file: %s", err.Error())
	}

	r := raspi.NewAdaptor()
	if err := r.Connect(); err != nil {
		log.Fatalf("error connecting to GPIO: %s", err.Error())
	}

	rtr := mux.NewRouter()

	switches := make(map[string]*TrainSwitch)
	for key, item := range opts.SwitchConfig {
		switches[key] = NewTrainSwitch(r,
			item.OnOffPin,
			item.DirectionPin)
	}

	ts := &TrainServer{
		switches: switches,
	}

	rtr.HandleFunc("/switch/{id}/toggle", ts.ToggleSwitch).Methods("GET")
	rtr.HandleFunc("/switch/{id}/status", ts.SwitchStatus).Methods("GET")
	rtr.HandleFunc("/switch/{id}/on", ts.TurnOn).Methods("GET")
	rtr.HandleFunc("/switch/{id}/off", ts.TurnOff).Methods("GET")

	srv := &http.Server{
		Handler: rtr,
		Addr:    fmt.Sprintf(":%d", opts.Port),
	}

	log.Printf("Listening on :%d...\n", opts.Port)
	log.Fatal(srv.ListenAndServe())
}
