package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("ws serving current directory on port 8080. Press CTRL+C to stop.")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
