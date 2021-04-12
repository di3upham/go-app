package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/urfave/cli"
)

var sampleapp *Sampleapp

func main() {
	sampleapp = NewSampleapp()
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{Name: "daemon", Usage: "run server", Action: daemon},
	}
	app.RunAndExitOnError()
}

func daemon(ctx *cli.Context) {
	go func() { http.ListenAndServe("localhost:6060", nil) }()
	sampleapp.ServeHTTP()
}
