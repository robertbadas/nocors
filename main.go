package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	port := flag.String("port", "5555", "Port to listen on")
	p := flag.String("p", "5555", "Port to listen on (shorthand)")
	flag.Parse()

	portToUse := *port
	if *p != "5555" {
		portToUse = *p
	}

	http.HandleFunc("/", handler)
	log.Printf("listening on :%s", portToUse)
	log.Fatal(http.ListenAndServe(":"+portToUse, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Get the target URL from the query parameter
	targetURL := r.URL.Query().Get("url")
	if targetURL == "" {
		http.Error(w, "url query parameter is missing", http.StatusBadRequest)
		return
	}
	var u, _ = url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(u)

	r.URL.Host = u.Host
	r.URL.Scheme = u.Scheme
	r.Host = u.Host
	// r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy.ServeHTTP(w, r)
}
