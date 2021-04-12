package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPPort string `default:":9000"`
}

type Sampleapp struct {
	*sync.Mutex
	Cf Config
}

func NewSampleapp() *Sampleapp {
	var cf Config
	envconfig.MustProcess("sampleapp", &cf)
	app := &Sampleapp{
		Mutex: &sync.Mutex{},
		Cf:    cf,
	}
	return app
}

func (app *Sampleapp) ServeHTTP() {
	fmt.Println("ServeHTTP at localhost" + app.Cf.HTTPPort)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong from sampleapp",
		})
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, struct{ Name string }{Name: "sampleapp"})
	})
	r.Run(app.Cf.HTTPPort)
}
