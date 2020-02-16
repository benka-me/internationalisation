package rpc_internl

import (
    "google.golang.org/grpc"
)

func ClientProvider (server string, options ...grpc.DialOption) (InternlClient, error) {
	conn, err := grpc.Dial(server + ":8030", options...)
	if err != nil {
		panic("Cannot dial")
	}

	return NewInternlClient(conn), nil
}

func ConnectThroughApi(server string, options ...grpc.DialOption) (InternlClient, error) {
	conn, err := grpc.Dial(server + ":8080", options...)
	if err != nil {
		panic("Cannot dial")
	}

	return NewInternlClient(conn), nil
}
