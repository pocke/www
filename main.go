package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/pocke/hlog"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/pflag"
)

func main() {
	var port int
	pflag.IntVarP(&port, "port", "p", 0, "TCP port number")
	pflag.Parse()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("http://127.0.0.1:%d", l.Addr().(*net.TCPAddr).Port)
	open.Run(url)
	fmt.Println(url)

	go reOpenner(url)

	http.Serve(l, hlog.Wrap(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "."+r.URL.Path)
	}))
}

func reOpenner(url string) {
	sc := bufio.NewScanner(os.Stdout)
	sc.Split(bufio.ScanLines)
	for sc.Scan() {

		t := sc.Text()
		if len(t) != 0 && t[0] == 'r' {
			open.Run(url)
		}
	}
}
