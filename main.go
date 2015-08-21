package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/skratchdot/open-golang/open"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("http://%s", l.Addr())
	open.Run(url)
	fmt.Println(url)

	http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+r.URL.Path)
	}))
}
