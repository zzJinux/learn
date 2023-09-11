package main

import (
	"flag"
	"net"
	"net/http"
	"log"
	"fmt"
)

var (
	fListenAddr = flag.String("listen", ":5678", "Listen address")
)

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		localAddr := r.Context().Value(http.LocalAddrContextKey).(net.Addr).String()
		fmt.Fprintf(w, "%s <= %s\n", localAddr, r.RemoteAddr)
	})

	log.Printf("Listening on %s ...", *fListenAddr)
	log.Fatal(http.ListenAndServe(*fListenAddr, nil))
}
