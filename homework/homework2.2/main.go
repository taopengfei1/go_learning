package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/setHeader", setHeader)
	http.HandleFunc("/returnVersion", returnVersion)
	http.HandleFunc("/accessLog", accessLog)
	http.HandleFunc("/health", health)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}

}

// return header
func setHeader(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		w.Header().Set(k, fmt.Sprintf("%s", v))
	}
}

// return version
func returnVersion(w http.ResponseWriter, r *http.Request) {
	VERSION := os.Getenv("VERSION")
	w.Header().Set("VERSION", VERSION)
}

// 打印访问日志
func accessLog(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Fatal(err)
	}
	err = flag.Set("v", "2")
	if err != nil {
		log.Fatal(err)
	}
	glog.V(2).Info("client ip: ", ip)
	glog.V(2).Info("http status code: 200")

}

// return 200
func health(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, fmt.Sprintf("200"))
	if err != nil {
		log.Fatal(err)
	}
}
