package cmd

import (
	"github.com/mfsiddique11/auth-server/pkg"
	"log"
	"net/http"
)

func RunServer() {
	println("Server is up on localhost:5000")
	router := pkg.Routes()
	log.Fatal(http.ListenAndServe(":5000", router))
}

