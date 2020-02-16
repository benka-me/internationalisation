package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/cell-user/go-pkg/user"
	rpcInternl "github.com/benka-me/internationalisation/go-pkg/rpc-internl"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Collection struct {
	Database string
	Messages string
}
type App struct {
	Client *mongo.Client
	Collection Collection
}

const port = ":8030"

var grpcServer *grpc.Server

func Server_2_0() {
	var err error
	app := &App {
		Collection: Collection{
			Messages:    "messages",
			Database: "internationalisation",
		},
	}

	ctx := context.Background()
	app.Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	grpcServer = grpc.NewServer()
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	{
		rpcInternl.RegisterInternlServer(grpcServer, app)
		reflection.Register(grpcServer)
	}
	fmt.Println("InternlServer Grpc 2.0 on port ", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
