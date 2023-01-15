package main

import (
	"io"
	"net/http"
	"os"
)

func setMuxHandler(mux *http.ServeMux) {
	mux.HandleFunc("/", testHandler)
}

func testHandler(rep http.ResponseWriter, req *http.Request) {

	io.WriteString(rep, "TestRouter")
}

func main() {

	listenAddr := os.Getenv("LISTEN_ADDR")
	listenPort := os.Getenv("LISTEN_PORT")

	if listenAddr == "" {

		listenAddr = "0.0.0.0"
		listenPort = "8080"
	}

	mux := http.NewServeMux()

	setMuxHandler(mux)
	http.ListenAndServe(listenAddr+":"+listenPort, mux)
}
