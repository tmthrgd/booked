package main

//go:generate go run -tags=dev internal/assets/generate.go
//go:generate go run -tags=dev internal/templates/generate.go

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":8080", "the address to listen on")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	h, err := open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("listening on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, h))
}
