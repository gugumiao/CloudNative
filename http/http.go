package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func healthz(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "200\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println(GetIP(req))
	for name, headers := range req.Header {
		for _, h := range headers {
			w.Header().Set(name, h)
			io.WriteString(w, fmt.Sprintf("%s=%s\n", name, h))
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
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", headers)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal(err)
	}
}
