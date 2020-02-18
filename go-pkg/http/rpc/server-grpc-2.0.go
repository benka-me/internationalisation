package rpc

import (
	"context"
	"fmt"
	//"github.com/benka-me/cell-user/go-pkg/user"
	rpcInternl "github.com/benka-me/internationalisation/go-pkg/rpc-internl"
	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type GrpcWebMiddleware struct {
	*grpcweb.WrappedGrpcServer
}

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
var router     *chi.Mux

func (m *GrpcWebMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "content-type, x-grpc-web")
		if m.IsAcceptableGrpcCorsRequest(r) || m.IsGrpcWebRequest(r) {
			m.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func NewGrpcWebMiddleware(grpcWeb *grpcweb.WrappedGrpcServer) *GrpcWebMiddleware {
	return &GrpcWebMiddleware{grpcWeb}
}
func Server_2_0() {
	router = chi.NewRouter()
	var err error
	app := &App {
		Collection: Collection{
			Database: "internationalisation",
			Messages: "messages",
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

	go func() {
		fmt.Println("InternlServer Grpc 2.0 on port ", port)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wrappedGrpc := grpcweb.WrapServer(grpcServer)
	router.Use(
		chiMiddleware.Logger,
		chiMiddleware.Recoverer,
		NewGrpcWebMiddleware(wrappedGrpc).Handler, // Must come before general CORS handling
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}).Handler,
	)
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {})

	if err := http.ListenAndServe(":8031", router); err != nil {
		grpclog.Fatalf("failed starting http2 Server: %v", err)
	}
}
