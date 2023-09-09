package handler

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

func pingCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "PONG!")
}

func registerPProfHandlers(r *http.ServeMux) {
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)
}

func RunServeHttp() {
	m := http.NewServeMux()
	m.HandleFunc("/ping", pingCheckHandler)
	registerPProfHandlers(m)
	log.Println("Server is run on port 3333")
	if err := http.ListenAndServe(":3333", m); err != nil {
		log.Fatal(err)
	}
}
