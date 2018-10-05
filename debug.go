package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"github.com/orvice/utils/env"
)

func pprof() {
	if env.Get("DEBUG") != "true"{
		return
	}
	log.Println(http.ListenAndServe(":6060", nil))
}
