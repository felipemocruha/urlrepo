package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/felipemocruha/urlrepo/urlrepo"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("[*] Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	mh := &ModelHandler{db: getDB()}
	pb.RegisterUrlServer(s, &server{modelHandler: mh})

	log.Printf("[*] Starting server at port %s", port)
	s.Serve(lis)
}
