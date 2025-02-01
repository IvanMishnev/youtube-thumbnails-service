package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/IvanMishnev/youtube-thumbnails-service/client-cli/helpers"
	pb "github.com/IvanMishnev/youtube-thumbnails-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	thumbnailsDir = "./thumbnails"
	address       = "localhost:3200"
)

func processThumbnail(c pb.YoutubeThumbnailsClient, url string) {
	request := &pb.GetThumbnailRequest{
		URL: url,
	}

	response, err := c.GetThumbnail(context.Background(), request)
	if err != nil {
		log.Printf("Error proccessing %s: %s", url, err.Error())
		return
	}

	if len(response.Error) > 0 {
		log.Printf("Error proccessing %s: %s", url, response.Error)
		return
	}

	img := response.Thumbnail

	videoID, err := helpers.GetVideoID(url)
	if err != nil {
		log.Println(err)
		return
	}

	fName := videoID + ".jpg"
	fPath := filepath.Join(thumbnailsDir, fName)

	err = os.WriteFile(fPath, img, os.ModePerm)
	if err != nil {
		log.Printf("Failed to write file: %s", err.Error())
		return
	}
}

func main() {
	err := runClient()
	if err != nil {
		log.Fatal(err)
	}
}

func runClient() error {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()

	c := pb.NewYoutubeThumbnailsClient(conn)

	urls, asyncFlag, err := helpers.ParseArgs()
	if err != nil {
		return err
	}

	err = os.MkdirAll(thumbnailsDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory for thumbnails")
	}

	if asyncFlag {
		log.Println("Async processing...")

		var wg sync.WaitGroup

		for _, url := range urls {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				processThumbnail(c, url)
			}(url)
		}
		wg.Wait()

	} else {
		log.Println("Simple processing...")

		for _, url := range urls {
			processThumbnail(c, url)
		}
	}

	log.Println("Processing completed")
	return nil
}
