package main

import (
	"fmt"
	"net/http"
)

func healthz(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "200")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println(GetIP(req))
	for name, headers := range req.Header {
		for _, h := range headers {
			w.Header().Set(name, h)
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func main() {
	http.HandleFunc("/", headers)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":12345", nil)
}
