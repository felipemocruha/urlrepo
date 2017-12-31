package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/felipemocruha/urlrepo/urlrepo"
)

const (
	address = "localhost:50051"
)

func addUrl(client pb.UrlClient, url *pb.UrlRequest) {
	resp, err := client.AddUrl(context.Background(), url)
	if err != nil {
		log.Fatalf("[*] Could not add URL: %+v", err)
	}
	if resp.Success {
		log.Printf("[*] URL added with id: %d", resp.Id)
	}
}

func getUrls(client pb.UrlClient) {
	stream, err := client.GetUrls(context.Background(), &pb.UrlFilter{})
	if err != nil {
		log.Fatalf("[*] Error on get urls: %v", err)
	}

	for {
		url, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("[*] %v.GetUrls(_) = _, %v", client, err)
		}
		log.Printf("[*] url: %v", url)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[*] Could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUrlClient(conn)

	url := &pb.UrlRequest{
		Url:   "https://google.com",
		Title: "google"}

	addUrl(client, url)
	getUrls(client)
}
