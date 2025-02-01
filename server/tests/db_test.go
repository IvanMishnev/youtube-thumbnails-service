package tests

import (
	"os"
	"testing"

	"github.com/IvanMishnev/youtube-thumbnails-service/server/cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDB(t *testing.T) {
	//временная бд для тестирования
	dbfFile := "temp.db"
	defer os.Remove(dbfFile)

	db, err := cache.NewCacheService(dbfFile)
	require.NoError(t, err)

	id := "some_id"
	img := []byte("some_image")

	//тест добавления записи в базу
	err = db.AddThumbnail(id, img)
	require.NoError(t, err, "Add test failed: expected no error")

	//тест получения записи из базы
	actualImg, err := db.GetThumbnail(id)
	require.NoError(t, err, "Get test failed: expected no error")
	assert.Equal(t, img, actualImg, "Get test failed: not equal")

	//обращение к несуществующей записи
	actualImg, err = db.GetThumbnail("010101")
	require.Error(t, err, "Get non-existent test failed: expected error")
	assert.Empty(t, actualImg, "Get non-existent test failed: expected empty image")

	//тест закрытия базы
	err = db.Close()
	require.NoError(t, err, "Close db test failed: expected no error")
}
