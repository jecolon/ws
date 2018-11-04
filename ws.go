package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.String("port", ":8080", "TCP port to listen on.")

func main() {
	flag.Parse()
	fmt.Printf("ws serving current directory on port %s. Press CTRL+C to stop.\n", *port)
	log.Fatal(http.ListenAndServe(*port, http.FileServer(http.Dir("."))))
}
