package app

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockEndpoint = func() int {
	return 1
}

var isRunning = false
var instance *Server

func setup() {
	if !isRunning {
		instance = MakeServer(mockEndpoint)
		instance.Start()
		isRunning = true
	}
}

func cleanup() {
	if instance != nil {
		instance.Close()
	}
}

func makeRequestWithResponse(path string) []byte {
	res := makeRequest(path)
	defer res.Body.Close()
	responseString, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return responseString
}

func makeRequest(path string) *http.Response {
	r, err := http.NewRequest(http.MethodGet, "http://localhost:3333"+path, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	return res
}

func TestServer(t *testing.T) {
	setup()
	responseString := makeRequestWithResponse("")
	assert.Equal(t, "1", string(responseString))

	badRequest := makeRequest("/bad")
	assert.Equal(t, 400, badRequest.StatusCode)

	goodRequest := makeRequestWithResponse("/1")
	assert.Equal(t, "1", string(goodRequest))

	goodRequest = makeRequestWithResponse("/17")
	assert.Equal(t, "0", string(goodRequest))
	cleanup()

}
