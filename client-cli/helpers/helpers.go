package helpers

import (
	"flag"
	"fmt"
	"strings"
)

func ParseArgs() ([]string, bool, error) {
	var urlsAll string
	var async bool

	flag.StringVar(&urlsAll, "URLs", "", "Usage: write video URLs separated by ',' without space")
	flag.BoolVar(&async, "async", false, "Use this argument for async processing")
	flag.Parse()

	if len(urlsAll) < 1 {
		return nil, false, fmt.Errorf("no URLs in arguments")
	}

	urls := strings.Split(urlsAll, ",")

	return urls, async, nil
}

func GetVideoID(URL string) (string, error) {
	var videoID string
	//Если передана ссылка формата www.youtube.com/watch?v=*videoID*
	if strings.Contains(URL, "youtube.com/watch?") {
		splited := strings.Split(URL, "=")

		videoID = splited[1]
		if len(videoID) < 1 {
			return "", fmt.Errorf("wrong video URL")
		}

		//Если ссылка формата youtu.be/*videoID*
	} else if strings.Contains(URL, "youtu.be") {
		index := strings.LastIndex(URL, "/")
		if index == -1 {
			return "", fmt.Errorf("wrong video URL")
		}
		videoID = URL[index+1:]
		if len(videoID) < 1 {
			return "", fmt.Errorf("wrong video URL")
		}
	} else {
		return "", fmt.Errorf("wrong video URL")
	}

	return videoID, nil
}
