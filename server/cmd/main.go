package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/IvanMishnev/youtube-thumbnails-service/proto"
	"github.com/IvanMishnev/youtube-thumbnails-service/server/cache"
	"github.com/IvanMishnev/youtube-thumbnails-service/server/constants"
	"github.com/IvanMishnev/youtube-thumbnails-service/server/helpers"
	"google.golang.org/grpc"
)

type ThumbnailsServer struct {
	pb.UnimplementedYoutubeThumbnailsServer
	CacheService *cache.CacheService
}

func NewThumbnailsServer(cacheService *cache.CacheService) *ThumbnailsServer {
	return &ThumbnailsServer{
		CacheService: cacheService,
	}
}

func (s *ThumbnailsServer) GetThumbnail(ctx context.Context, req *pb.GetThumbnailRequest) (*pb.GetThumbnailResponse, error) {
	url := req.URL

	videoID, err := helpers.GetVideoID(url)
	if err != nil {
		log.Printf("failed to get video id for %s: %s", url, err.Error())
		return &pb.GetThumbnailResponse{Error: err.Error()}, nil
	}

	//Поиск изображения в кэше
	img, err := s.CacheService.GetThumbnail(videoID)
	//Если изображение отсутствует в кэше, загружаем и добавляем
	if err == cache.ErrNotFound {
		log.Printf("image for %s not found in cache. Downloading...", videoID)

		img, err = helpers.DownloadThumbnail(videoID)
		if err != nil {
			log.Printf("failed to download thumbnail for %s", videoID)
			return &pb.GetThumbnailResponse{Error: err.Error()}, nil
		}

		err = s.CacheService.AddThumbnail(videoID, img)
		if err != nil {
			log.Println(err)
		}
	} else if err != nil {
		//В случае других ошибок
		log.Printf("failed to get image from cache: %s", err.Error())

		img, err = helpers.DownloadThumbnail(videoID)
		if err != nil {
			log.Printf("failed to download thumbnail for %s", videoID)
			return &pb.GetThumbnailResponse{Error: err.Error()}, nil
		}
	}

	response := pb.GetThumbnailResponse{
		Thumbnail: img,
	}
	return &response, nil
}

func main() {
	err := startServer()
	if err != nil {
		log.Fatalf("failed to start server: %s", err.Error())
	}
}

func startServer() error {
	listen, err := net.Listen("tcp", constants.Address)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	cacheService, err := cache.NewCacheService(constants.CachePath)
	defer cacheService.Close()

	if err != nil {
		return fmt.Errorf("failed to create cache: %w", err)
	}

	s := grpc.NewServer()
	thumbnailsSrv := NewThumbnailsServer(cacheService)
	pb.RegisterYoutubeThumbnailsServer(s, thumbnailsSrv)

	err = s.Serve(listen)
	if err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
