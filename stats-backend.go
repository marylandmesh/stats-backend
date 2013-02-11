package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

// Flags
var (
	fIface = flag.String("iface", "", "interface to listen on")
	fPort  = flag.String("port", "8080", "port to listen on")
	fRoot  = flag.String("root", "/", "url that should be used as root for the listener")

	fLog = flag.Bool("log", true, "whether to enable logging")
)

var (
	l *log.Logger
)

func main() {
	flag.Parse()

	if *fLog {
		l = log.New(os.Stdout, "", log.Ldate|log.Ltime)
	} else {
		l = log.New(ioutil.Discard, "", 0)
	}

	err := serve(*fIface, *fPort)
	if err != nil {
		l.Fatalln(err)
	}
}
