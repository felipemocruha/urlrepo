package main

import (
	"github.com/asdine/storm"
	pb "github.com/felipemocruha/urlrepo/urlrepo"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"log"
	"strings"
	"time"
)

type UrlModel struct {
	Id      string `storm:"id"`
	Url     string `storm:"unique"`
	Title   string
	AddedAt int
}

type ModelHandler struct {
	db *storm.DB
}

type server struct {
	modelHandler *ModelHandler
}

func (s *server) AddUrl(ctx context.Context, in *pb.UrlRequest) (*pb.UrlResponse, error) {
	err := s.modelHandler.createUrl(in)
	if err != nil {
		return &pb.UrlResponse{Id: "", Success: false}, err
	}
	return &pb.UrlResponse{Id: in.Id, Success: true}, nil
}

func (s *server) GetUrls(filter *pb.UrlFilter, stream pb.Url_GetUrlsServer) error {
	urls := s.modelHandler.listUrls()
	for _, url := range urls {
		request := modelToRequest(&url)
		if err := stream.Send(&request); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) GetUrl(ctx context.Context, in *pb.UrlFilter) (*pb.UrlRequest, error) {
	resp := s.modelHandler.selectUrl(in.Id)
	req := modelToRequest(&resp)
	return &req, nil
}

func (s *server) RemoveUrl(ctx context.Context, in *pb.UrlFilter) (*pb.UrlResponse, error) {
	err := s.modelHandler.removeUrl(in.Id)
	if err != nil {
		return &pb.UrlResponse{Id: "", Success: false}, nil
	}
	return &pb.UrlResponse{Id: in.Id, Success: true}, nil
}

func getDB() *storm.DB {
	dbPath := getenv("DB_PATH", "db")
	db, err := storm.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (mh *ModelHandler) createUrl(url *pb.UrlRequest) error {
	id := strings.Replace(uuid.NewV4().String(), "-", "", -1)
	url.Id = id
	url.Title = fetchUrlTitle(url.Url)
	url.AddedAt = int32(time.Now().Unix())
	model := requestToModel(url)
	err := mh.db.Save(&model)

	return err
}

func (mh *ModelHandler) listUrls() []UrlModel {
	var urls []UrlModel
	err := mh.db.All(&urls)
	if err != nil {
		panic(err)
	}

	return urls
}

func (mh *ModelHandler) selectUrl(id string) UrlModel {
	url := UrlModel{}
	err := mh.db.One("Id", id, &url)
	if err != nil {
		panic(err)
	}

	return url
}

func (mh *ModelHandler) removeUrl(id string) error {
	url := UrlModel{}
	err := mh.db.One("Id", id, &url)
	if err != nil {
		return err
	}

	err = mh.db.DeleteStruct(&url)
	return err
}

func modelToRequest(url *UrlModel) pb.UrlRequest {
	return pb.UrlRequest{
		Id:      url.Id,
		Url:     url.Url,
		Title:   url.Title,
		AddedAt: int32(url.AddedAt)}
}

func requestToModel(url *pb.UrlRequest) UrlModel {
	return UrlModel{
		Id:      url.Id,
		Url:     url.Url,
		Title:   url.Title,
		AddedAt: int(url.AddedAt)}
}
