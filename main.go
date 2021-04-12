package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"

	pb "git.local/go-app/model"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

var sampleapp *Sampleapp

func main() {
	sampleapp = NewSampleapp()
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{Name: "daemon", Usage: "run server", Action: daemon},
		{Name: "grpchello", Usage: "run grpc client", Action: grpchello},
		{Name: "grpcordercreate", Action: grpcordercreate},
		{Name: "grpcorderread", Action: grpcorderread},
	}
	app.RunAndExitOnError()
}

func daemon(ctx *cli.Context) {
	go func() { http.ListenAndServe("localhost:6060", nil) }()
	sampleapp.BatchAsync()
	go sampleapp.ServeGrpc()
	sampleapp.ServeHTTP()
}

func grpchello(clictx *cli.Context) {
	// Set up a connection to the server.
	target := "localhost" + sampleapp.Cf.GrpcPort
	fmt.Println(target)
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("2")
	defer conn.Close()
	c := pb.NewSampleAPIClient(conn)

	// Contact the server and print out its response.
	name := ""
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("1")
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func grpcordercreate(clictx *cli.Context) {
	// Set up a connection to the server.
	target := "localhost" + sampleapp.Cf.GrpcPort
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleAPIClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	productUrl := strings.TrimSpace(clictx.Args().Get(0))
	orderStatus := strings.TrimSpace(clictx.Args().Get(1))
	r, err := c.CreateOrder(ctx, &pb.Order{ProductUrl: productUrl, Status: orderStatus})
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	if r != nil {
		log.Printf("Order: %s", r)
	}
}

func grpcorderread(clictx *cli.Context) {
	// Set up a connection to the server.
	target := "localhost" + sampleapp.Cf.GrpcPort
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleAPIClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	orderId := strings.TrimSpace(clictx.Args().Get(0))
	r, err := c.ReadOrder(ctx, &pb.Id{Id: orderId})
	if err != nil {
		log.Fatalf("Err: %v", err)
	}
	if r != nil {
		log.Printf("Order: %s", r)
	}
}
