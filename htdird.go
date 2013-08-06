package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Version stores current version number.
const Version = "0.0.1"

// Config structure represents app's configuration.
type Config struct {
	Addr  string // Listen address.
	Dir   string // Directory meant to be served.
	Debug bool   // Enable debug mode.
}

// conf stores global configuration.
var conf = new(Config)

// showVersion when set to true makes the app to print current version and exit.
var showVersion = false

// debugf is a default debugging function. If debug mode is enabled it will be
// replaced with log.Printf function.
var debugf = func(f string, a ...interface{}) {}

func init() {
	log.SetFlags(0)

	flag.BoolVar(&conf.Debug, "d", false, "enable debug mode")
	flag.BoolVar(&showVersion, "v", false, "display version number and exit")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 2 {
		conf.Addr = flag.Arg(0)
		conf.Dir = flag.Arg(1)
	}
}

func main() {
	// If ran with -v flag, display version number and exit.
	if showVersion {
		fmt.Println(Version)
		return
	}

	// Debug mode enabled? Set log.Printf as debug function.
	if conf.Debug {
		debugf = log.Printf
		debugf("D -- debug mode enabled")
	}

	// Validate address and directory.
	if conf.Addr == "" {
		log.Fatalf("E -- missing address")
	}
	if conf.Dir == "" {
		log.Fatalf("E -- missing directory")
	}

	debugf("D -- provided configuration: %+v", conf)

	// Serve directory...
	log.Printf("I -- about to start serving '%s' on '%s'", conf.Dir, conf.Addr)
	fileServer := http.FileServer(http.Dir(conf.Dir))
	log.Fatalf("E -- %v", http.ListenAndServe(conf.Addr, serve(fileServer)))
}

// usage displays application help screen.
func usage() {
	fmt.Print(strings.TrimLeft(Help, "\n"))
}

// ResponseWriter is a tiny wrapper around standard response writer that adds
// possibility to retrieve the status after it's set.
type ResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.Status = status
}

// serve takes the file server handler and wraps it up with logging
// and extra security improvements.
func serve(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &ResponseWriter{w, 200}

		// Allow only GET methods!
		if r.Method != "GET" {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			rw.Write([]byte("405 method not allowed"))
			return
		}

		handler.ServeHTTP(rw, r)
		log.Printf("I -- %s - %d - %s %s", r.RemoteAddr, rw.Status, r.Method, r.URL)
	})
}

const Help = `
Usage:

  htdird [-d] [-h] [-v] ADDR DIR

Start HTTP server on address ADDR that will serve static files from
directory DIR.

Options:

  -d  Enables debug mode.
  -h  Shows help screen.
  -v  Print version number and exit.

For more information check 'man htdird'.
`
