package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"gopkg.in/alecthomas/kingpin.v2"
)

const VERSION = "1.1.0"

var (
	app      = kingpin.New("servedir", "Serve a directory over HTTP(S)")
	addr     = app.Flag("addr", "Address to listen on").Short('a').Default("127.0.0.1:8080").String()
	dir      = app.Flag("dir", "Directory to serve").Short('d').Default(".").String()
	certFile = app.Flag("cert", "TLS certificate to use (optional)").String()
	keyFile  = app.Flag("key", "TLS certificate private key to use when using --cert").String()
)

func main() {
	app.Version(VERSION)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	fileServer := http.FileServer(http.Dir(*dir))
	http.Handle("/", handlers.CombinedLoggingHandler(os.Stdout, fileServer))

	if *certFile == "" && *keyFile == "" {
		log.Printf("Starting HTTP fileserver on %s to serve directory %s", *addr, *dir)
		log.Fatal(http.ListenAndServe(*addr, nil))
	} else {
		log.Printf("Starting HTTPS fileserver on %s to serve directory %s", *addr, *dir)
		log.Fatal(http.ListenAndServeTLS(*addr, *certFile, *keyFile, nil))
	}
}
