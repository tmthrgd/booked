package main

//go:generate go run -tags=dev internal/assets/generate.go
//go:generate go run -tags=dev internal/templates/generate.go

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"go.tmthrgd.dev/booked/browser"
)

func init() {
	log.SetFlags(0)

	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "usage: %s [flags] facebook-json.zip\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	addr := flag.String("addr", ":8080", "the address to listen on")
	noLaunch := flag.Bool("no-launch", false, "skip launching a browser window")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	h, err := open(flag.Args()...)
	if err != nil {
		log.Fatal(err)
	}

	if !*noLaunch {
		go launch(*addr)
	}

	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, h))
}

func launch(addr string) {
	u := &url.URL{
		Scheme: "http",
		Host:   addr,
	}
	if u.Hostname() == "" {
		u.Host = net.JoinHostPort("localhost", u.Port())
	}

	if !browser.Open(u.String()) {
		log.Printf("navigate to %s in your browser", u)
	}
}
