package main

import (
	"flag"
	"os"
	"testing"

	"github.com/IvanMishnev/youtube-thumbnails-service/client-cli/helpers"
	"github.com/stretchr/testify/assert"
)

func TestParseArgs(t *testing.T) {
	//сохраним оригинальные значения flag и восстановим их после окончания теста
	originalArgs := flag.CommandLine.Args()
	defer func() { flag.CommandLine.Parse(originalArgs) }()

	//case 1: Valid URLs and no async flag
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	os.Args = []string{"test", "--URLs", "https://www.youtube.com/watch?v=N76ErzOdk9g,https://youtu.be/ujChUYkPvec"}
	flag.CommandLine.Parse(os.Args[1:])

	urls, async, err := helpers.ParseArgs()
	assert.NoError(t, err, "Case 1 test failed: expected no error")

	expectedUrls := []string{"https://www.youtube.com/watch?v=N76ErzOdk9g", "https://youtu.be/ujChUYkPvec"}
	assert.Equal(t, expectedUrls, urls, "Case 1 test failed: not equal URLs")
	assert.False(t, async, "Case 1 test failed: expected async is false")

	//case 2: Valid URLs and async flag
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	os.Args = []string{"test", "--URLs", "https://www.youtube.com/watch?v=N76ErzOdk9g,https://youtu.be/ujChUYkPvec", "--async"}
	flag.CommandLine.Parse(os.Args[1:])

	urls, async, err = helpers.ParseArgs()
	assert.NoError(t, err, "Case 2 test failed: Expected no error")

	expectedUrls = []string{"https://www.youtube.com/watch?v=N76ErzOdk9g", "https://youtu.be/ujChUYkPvec"}
	assert.Equal(t, expectedUrls, urls, "Case 2 test failed: not equal URLs")
	assert.True(t, async, "Case 2 test failed: expected async is true")

	//case 3: No URLs
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	os.Args = []string{"test"}
	flag.CommandLine.Parse(os.Args[1:])

	_, _, err = helpers.ParseArgs()
	assert.Error(t, err, "Case 3 test failed: expected error")

	//case 4: Empty URLs string
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	os.Args = []string{"test", "--URLs", ""}
	flag.CommandLine.Parse(os.Args[1:])

	_, _, err = helpers.ParseArgs()
	assert.Error(t, err, "Case 4 test failed: expected error")
}
