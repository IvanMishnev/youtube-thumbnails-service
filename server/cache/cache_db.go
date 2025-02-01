package cache

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type CacheService struct {
	db *sql.DB
}

var ErrNotFound = fmt.Errorf("not found in cache")

func NewCacheService(dbPath string) (*CacheService, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS thumbnails (
				video_id VARCHAR(256) PRIMARY KEY,
				image BLOB
				)`)
	if err != nil {
		return nil, fmt.Errorf("failed to execute init query: %w", err)
	}

	log.Println("Connected to cache database")
	return &CacheService{
		db: db,
	}, nil
}

func (c *CacheService) AddThumbnail(videoID string, thumbnailImage []byte) error {
	_, err := c.db.Exec(`INSERT INTO thumbnails (video_id, image) VALUES (?, ?)`, videoID, thumbnailImage)
	if err != nil {
		return fmt.Errorf("failed to add thumbnail image to cache: %w", err)
	}
	log.Printf("Thumbnail for %s added to cache", videoID)

	return nil
}

func (c *CacheService) GetThumbnail(videoId string) ([]byte, error) {
	var image []byte
	row := c.db.QueryRow(`SELECT image FROM thumbnails WHERE video_id = ?`, videoId)

	err := row.Scan(&image)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, fmt.Errorf("failed to get thumbnail from cache: %w", err)
	}

	log.Printf("image for %s downloaded from cache succesfully", videoId)
	return image, nil
}

func (c *CacheService) Close() error {
	err := c.db.Close()
	if err != nil {
		return err
	}

	return nil
}
