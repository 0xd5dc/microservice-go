package main

import (
	"fmt"
	"github.com/unrolled/render"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	formatter = render.New(render.Options{

		IndentJSON: true,
	})
)

func TestCreateEvent(t *testing.T) {
	client := &http.Client{}

	server := httptest.NewServer(

		http.HandlerFunc(createEventHandler(formatter)))

	defer server.Close()
	fmt.Print(client)
}
