package main

import (
	"flag"
	"net/http"
)

var (
	responseText   string
	networkAddress string
)

func parseFlags() {
	flag.StringVar(&responseText, "text", "ok", "Text to reply")
	flag.StringVar(&networkAddress, "bind", ":8000",
		"Network address for listening")
	flag.Parse()
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	h := w.Header()
	h.Set("Cache-Control", "no-cache")
	h.Set("Pragma", "no-cache")
	h.Set("Content-Type", "text/plain")
	w.Write([]byte(responseText))
}

func main() {
	parseFlags()
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(networkAddress, nil)
}
