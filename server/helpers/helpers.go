package helpers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

var ErrWrongURL = fmt.Errorf("wrong video URL")

// Функция возвращает id видео, полученный из полного URL
func GetVideoID(URL string) (string, error) {
	var videoID string
	//Если передана ссылка формата www.youtube.com/watch?v=*videoID*
	if strings.Contains(URL, "youtube.com/watch?") {
		splited := strings.Split(URL, "=")

		videoID = splited[1]
		if len(videoID) < 1 {
			return "", ErrWrongURL
		}

		//Если ссылка формата youtu.be/*videoID*
	} else if strings.Contains(URL, "youtu.be") {
		index := strings.LastIndex(URL, "/")
		if index == -1 {
			return "", ErrWrongURL
		}
		videoID = URL[index+1:]
		if len(videoID) < 1 {
			return "", ErrWrongURL
		}

	} else {
		return "", ErrWrongURL
	}
	return videoID, nil
}

func DownloadThumbnail(videoID string) ([]byte, error) {
	url := "https://img.youtube.com/vi/" + videoID + "/hqdefault.jpg"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	img, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}
