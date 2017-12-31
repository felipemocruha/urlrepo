package main

import (
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/felipemocruha/urlrepo/urlrepo"
	"os"
)

const (
	address = "localhost:50051"
)

func addUrl(client pb.UrlClient, url *pb.UrlRequest) {
	_, err := client.AddUrl(context.Background(), url)
	if err != nil {
		log.Fatalf("[*] Could not add URL: %v", err)
	}
}

func getUrl(client pb.UrlClient, filter *pb.UrlFilter) {
	resp, err := client.GetUrl(context.Background(), filter)
	if err != nil {
		log.Fatalf("[*] Could not get URL: %v", err)
	}
	log.Println(resp)
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
		log.Printf("[*] url: %+v", url)
	}
}

func deleteUrl(client pb.UrlClient, filter *pb.UrlFilter) {
	_, err := client.RemoveUrl(context.Background(), filter)
	if err != nil {
		log.Fatalf("[*] Could not remove URL: %v", err)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("[*] Could not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewUrlClient(conn)

	if os.Args[1] == "add" {
		url := &pb.UrlRequest{Url: os.Args[2]}
		addUrl(client, url)

	} else if os.Args[1] == "get" {
		filter := &pb.UrlFilter{Id: os.Args[2]}
		getUrl(client, filter)

	} else if os.Args[1] == "rm" {
		filter := &pb.UrlFilter{Id: os.Args[2]}
		deleteUrl(client, filter)

	} else if os.Args[1] == "list" {
		getUrls(client)
	}

}
