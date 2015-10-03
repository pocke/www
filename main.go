package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/pocke/hlog"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("http://127.0.0.1:%d", l.Addr().(*net.TCPAddr).Port)
	open.Run(url)
	fmt.Println(url)

	http.Serve(l, hlog.Wrap(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+r.URL.Path)
	}))
}
