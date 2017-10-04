package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

const defaultPort = 8080

func main() {
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}
	listenAddr := fmt.Sprintf(":%d", port)
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/metrics", metricsHandler)
	http.ListenAndServe(listenAddr, nil)
}

func getPort() (uint16, error) {
	env := os.Getenv("OPENSHIFT_PROMETHEUS_TARGET_EXAMPLE")
	if env == "" {
		return defaultPort, nil
	}
	n, err := strconv.ParseUint(env, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("OPENSHIFT_PROMETHEUS_TARGET_EXAMPLE environment variable invalid value: %v", err)
	}
	return uint16(n), nil
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	alsoDumpBody := true
	dump, err := httputil.DumpRequest(r, alsoDumpBody)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		log.Printf("default handler: cannot process request: %v: cannot dump request: %v\n", r.URL.Path, err)
		return
	}

	if _, err := fmt.Fprintf(w, "Dump of your request:\n\n%s", dump); err != nil {
		log.Printf("default handler: cannot process request: %v: cannot write response: %v\n", r.URL.Path, err)
	}
	log.Println("default handler: successfully processed request:", r.URL.Path)
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprint(w, "TODO add some metrics\n"); err != nil {
		log.Printf("metrics handler: cannot process request: %v: cannot write response: %v\n", r.URL.Path, err)
	}
	log.Println("mstrics handler: successfully processed request:", r.URL.Path)
}
