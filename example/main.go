package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		fmt.Fprintf(w, "Your ip:"+ip+", f11111_00")
	})

	fmt.Println("Start listening on port:", 9090)
	http.ListenAndServe(":9090", nil)
}

