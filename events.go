package main

import (
	"github.com/unrolled/render"
	"net/http"
)

func createEventHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		w.Header().Add("Location", "shit")
		formatter.JSON(w,
			http.StatusCreated,
			struct{ Test string }{"This is a test"})

	}
}
