package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/pocke/hlog"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/pflag"
)

func main() {
	conf, err := loadConfigFile()
	if err != nil {
		panic(err)
	}

	var port int
	var binding string
	fs := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
	fs.IntVarP(&port, "port", "p", 0, "TCP port number")
	fs.StringVarP(&binding, "binding", "b", "localhost", "Bind www to the specified IP.")
	fs.Parse(append(conf, os.Args...))

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", binding, port))
	if err != nil {
		panic(err)
	}
	url := fmt.Sprintf("http://127.0.0.1:%d", l.Addr().(*net.TCPAddr).Port)
	open.Run(url)
	fmt.Println(url)

	go reOpenner(url)

	http.Serve(l, hlog.Wrap(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-store")
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

func loadConfigFile() ([]string, error) {
	path := "./.www"
	if !fileExist(path) {
		return []string{}, nil
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := strings.Trim(string(b), "\n")

	return strings.Split(s, " "), nil
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
