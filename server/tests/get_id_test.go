package tests

import (
	"testing"

	"github.com/IvanMishnev/youtube-thumbnails-service/server/helpers"
	"github.com/stretchr/testify/assert"
)

type want struct {
	id  string
	err error
}

var ErrWrongURL = helpers.ErrWrongURL

func TestGetVideoID(t *testing.T) {
	videos := map[string]want{
		"https://www.youtube.com/watch?v=CWeXOm49kE0":    {"CWeXOm49kE0", nil},
		"https://www.youtube.com/watch?v=WqEweV0eScg":    {"WqEweV0eScg", nil},
		"https://youtu.be/ujChUYkPvec":                   {"ujChUYkPvec", nil},
		"https://youtu.be/-488UORrfJ0":                   {"-488UORrfJ0", nil},
		"https://www.youtube.com/waasatch?v=N76ErzOdk9g": {"", ErrWrongURL},
		"https://www.youtube.com/watch?v=":               {"", ErrWrongURL},
		"https://youtu.b/":                               {"", ErrWrongURL},
		"random":                                         {"", ErrWrongURL},
	}

	for k, v := range videos {
		id, err := helpers.GetVideoID(k)

		assert.ErrorIs(t, err, v.err)
		assert.Equal(t, id, v.id)
	}
}
