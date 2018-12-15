package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vrischmann/envconfig"
	"gobot.io/x/gobot/platforms/raspi"
)

type Options struct {
	Port int `envconfig:"default=8080"`
}

type TrainServer struct {
	rtr      *mux.Router
	switches map[string]*TrainSwitch
}

func (ts *TrainServer) ToggleSwitch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sw, ok := ts.switches[vars["id"]]
	if ok {
		sw.Toggle()
		if sw.Status() {
			fmt.Fprint(w, "1")
		} else {
			fmt.Fprint(w, "0")
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func (ts *TrainServer) TurnOn(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sw, ok := ts.switches[vars["id"]]
	if ok {
		sw.On()
		fmt.Fprint(w, "1")
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func (ts *TrainServer) TurnOff(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sw, ok := ts.switches[vars["id"]]
	if ok {
		sw.Off()
		fmt.Fprint(w, "0")
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

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
	var opts Options
	if err := envconfig.Init(&opts); err != nil {
		log.Fatal(err.Error())
	}
	r := raspi.NewAdaptor()
	if err := r.Connect(); err != nil {
		log.Fatal(err.Error())
	}

	rtr := mux.NewRouter()

	ts := &TrainServer{
		rtr: rtr,
		switches: map[string]*TrainSwitch{
			"1": NewTrainSwitch(r, "35", "37"),
			"2": NewTrainSwitch(r, "11", "13"),
			"3": NewTrainSwitch(r, "16", "18"),
			"4": NewTrainSwitch(r, "8", "10"),
		},
	}

	rtr.HandleFunc("/toggle/{id}", ts.ToggleSwitch).Methods("GET")
	rtr.HandleFunc("/status/{id}", ts.SwitchStatus).Methods("GET")
	rtr.HandleFunc("/on/{id}", ts.TurnOn).Methods("GET")
	rtr.HandleFunc("/off/{id}", ts.TurnOff).Methods("GET")

	srv := &http.Server{
		Handler: rtr,
		Addr:    fmt.Sprintf(":%d", opts.Port),
	}

	log.Printf("Listening on :%d...\n", opts.Port)
	log.Fatal(srv.ListenAndServe())
}
