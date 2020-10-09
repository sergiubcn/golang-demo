package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	booksJSON := `[{"id":"1","isbn":"111","title":"Book One","author":{"firstname":"John","lastname":"Doe"}},{"id":"2","isbn":"222","title":"Book Two","author":{"firstname":"Steve","lastname":"Smith"}}]`
	response, err := http.Get("http://localhost:9090/api/books")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(200, response.StatusCode, "Response code should be 200 (OK)")
	assert.Equal(booksJSON, strings.ReplaceAll(string(body), "\n", ""), "The response body should be equal to booksJSON")
}

func TestGetBook(t *testing.T) {
	bookJSON := `{"id":"2","isbn":"222","title":"Book Two","author":{"firstname":"Steve","lastname":"Smith"}}`
	response, err := http.Get("http://localhost:9090/api/books/2")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	assert := assert.New(t)
	assert.Nil(err)
	assert.Equal(200, response.StatusCode, "Response code should be 200 (OK)")
	assert.Equal(bookJSON, strings.ReplaceAll(string(body), "\n", ""), "The response body should be equal to booksJSON")
}
