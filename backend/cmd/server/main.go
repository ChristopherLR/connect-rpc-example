package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"connectrpc.com/connect"
	"github.com/rs/cors"

	greetv1 "example.com/connect-example/greet/v1"
	"example.com/connect-example/greet/v1/greetv1connect"
)

type GreetServer struct {
	greetv1connect.UnimplementedGreetServiceHandler
}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Printf("Got request: %s", req.Msg.Name)
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	return res, nil
}

func (s *GreetServer) StreamGreetings(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
	stream *connect.ServerStream[greetv1.GreetResponse],
) error {
	log.Printf("Got stream request: %s", req.Msg.Name)
	for i := 0; i < 5; i++ {
		if err := stream.Send(&greetv1.GreetResponse{
			Greeting: fmt.Sprintf("Hello #%d, %s!", i+1, req.Msg.Name),
		}); err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)

	// CORS setup
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // In production, replace with specific origin
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Content-Encoding",
			"Connect-Accept-Encoding",
			"Grpc-Timeout",
			"X-Grpc-Web",
			"X-User-Agent",
		},
		ExposedHeaders: []string{
			"Content-Encoding",
			"Connect-Content-Encoding",
			"Grpc-Status",
			"Grpc-Message",
		},
	})

	fmt.Println("Starting server on localhost:8080")
	if err := http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(corsHandler.Handler(mux), &http2.Server{}),
	); err != nil {
		log.Fatalf("Listen failed: %v", err)
	}
}
