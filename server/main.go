package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"

	counterv1 "comp/server/gen/counter/v1"
	"comp/server/gen/counter/v1/counterv1connect"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var counter atomic.Int64

type CounterServer struct{}

func (s *CounterServer) Increment(
	_ context.Context,
	_ *connect.Request[counterv1.IncrementRequest],
) (*connect.Response[counterv1.IncrementResponse], error) {
	newVal := counter.Add(1)
	return connect.NewResponse(&counterv1.IncrementResponse{
		Value: newVal,
	}), nil
}

func (s *CounterServer) GetValue(
	_ context.Context,
	_ *connect.Request[counterv1.GetValueRequest],
) (*connect.Response[counterv1.GetValueResponse], error) {
	return connect.NewResponse(&counterv1.GetValueResponse{
		Value: counter.Load(),
	}), nil
}

func main() {
	mux := http.NewServeMux()
	path, handler := counterv1connect.NewCounterServiceHandler(&CounterServer{})
	mux.Handle(path, handler)

	addr := "localhost:8080"
	fmt.Printf("Server listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	))
}
