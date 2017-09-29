package main

import (
	"fmt"
	"log"
	"net/http"
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
	http.HandleFunc("/", handler)
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

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("request:", r.URL.Path)
	fmt.Fprintf(w, "%s", r.URL.Path[1:])
}
