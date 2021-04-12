package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"

	pb "git.local/go-app/model"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type Config struct {
	HTTPPort string `default:":9000"`
	GrpcPort string `default:":50051"`
}

type Sampleapp struct {
	*sync.Mutex
	Cf        Config
	stopchan  chan struct{} // signal to stop scheduling
	isrunning bool
	pb.UnimplementedSampleAPIServer
	ordermgr *Ordermgr
}

func NewSampleapp() *Sampleapp {
	var cf Config
	envconfig.MustProcess("sampleapp", &cf)
	app := &Sampleapp{
		Mutex:    &sync.Mutex{},
		Cf:       cf,
		stopchan: make(chan struct{}),
		ordermgr: NewOrdermgr(NewDB()),
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
	r.GET("/orders/:id", func(c *gin.Context) {
		out, err := app.ordermgr.ReadOrder(c.Param("id"))
		if err != nil {
			// TODO errconv
			c.JSON(500, err)
			return
		}
		c.JSON(200, out)
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

func (app *Sampleapp) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (app *Sampleapp) ServeGrpc() {
	fmt.Println("ServeGrpc at localhost" + app.Cf.GrpcPort)
	lis, err := net.Listen("tcp", app.Cf.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSampleAPIServer(s, app)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (app *Sampleapp) CreateOrder(ctx context.Context, in *pb.Order) (*pb.Order, error) {
	// TODO perm
	out, err := app.ordermgr.CreateOrder(in)
	if err != nil {
		// TODO errconv
		return nil, err
	}
	return out, err
}

func (app *Sampleapp) DeleteOrder(context.Context, *pb.Id) (*pb.Empty, error) {
	// TODO
	return nil, nil
}

func (app *Sampleapp) ListOrders(context.Context, *pb.Id) (*pb.Orders, error) {
	// TODO
	return nil, nil
}

func (app *Sampleapp) UpdateOrder(context.Context, *pb.Order) (*pb.Order, error) {
	// TODO
	return nil, nil
}

func (app *Sampleapp) ReadOrder(ctx context.Context, in *pb.Id) (*pb.Order, error) {
	// TODO perm
	out, err := app.ordermgr.ReadOrder(in.GetId())
	if err != nil {
		// TODO errconv
		return nil, err
	}
	return out, err
}
