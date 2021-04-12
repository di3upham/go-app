package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPPort string `default:":9000"`
}

type Sampleapp struct {
	*sync.Mutex
	Cf        Config
	stopchan  chan struct{} // signal to stop scheduling
	isrunning bool
}

func NewSampleapp() *Sampleapp {
	var cf Config
	envconfig.MustProcess("sampleapp", &cf)
	app := &Sampleapp{
		Mutex:    &sync.Mutex{},
		Cf:       cf,
		stopchan: make(chan struct{}),
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

func (app *Sampleapp) BatchAsync() chan struct{} {
	fmt.Println("BATCHASYNC-STARTTT")
	app.Lock()
	defer app.Unlock()
	if app.isrunning {
		return app.stopchan
	}
	app.isrunning = true
	ticker := time.NewTicker(1 * time.Second)
	var tickerCount int64
	go func() {
		for {
			select {
			case <-ticker.C:
				atomic.AddInt64(&tickerCount, 1)
				if tickerCount%5 == 0 {
					go func() { fmt.Println("TODO", tickerCount) }()
				}
			case <-app.stopchan:
				ticker.Stop()
				app.isrunning = false
				fmt.Println("BATCHASYNC-ENDDD")
				return
			}
		}
	}()

	return app.stopchan
}
