package main

import (
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
	if err := Main(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var version = "master"

func Main(args []string) error {
	conf, err := loadConfigFile()
	if err != nil {
		return err
	}

	var port int
	var binding string
	var noBrowser bool
	var displayVersion bool
	fs := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	fs.IntVarP(&port, "port", "p", 0, "TCP port number")
	fs.StringVarP(&binding, "binding", "b", "localhost", "Bind www to the specified IP.")
	fs.BoolVarP(&noBrowser, "no-browser", "n", false, "Do not open a browser.")
	fs.BoolVarP(&displayVersion, "version", "v", false, "Display version")
	err = fs.Parse(append(conf, os.Args...))
	if err != nil {
		if err == pflag.ErrHelp {
			return nil
		} else {
			return err
		}
	}

	if displayVersion {
		fmt.Println(version)
		return nil
	}

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", binding, port))
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://127.0.0.1:%d", l.Addr().(*net.TCPAddr).Port)
	if !noBrowser {
		if err := open.Run(url); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
	fmt.Println(url)

	return http.Serve(l, hlog.Wrap(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-store")
		http.ServeFile(w, r, "."+r.URL.Path)
	}))
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
