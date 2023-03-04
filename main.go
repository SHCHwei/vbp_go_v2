package main

import (
	"log"
	"net/http"
	_ "vbp/cmd/vbpServer/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":8090", nil))
}
