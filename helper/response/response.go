package response

import (
	"log"
	"net/http"
	"strconv"
)

type HttpResponse interface {
	WriteResponse(w http.ResponseWriter)
}

type responseCloser func()

type BasicResponse struct {
	close       responseCloser
	Body        []byte
	StatusCode  int
	ContentType string
}

func (b *BasicResponse) WriteResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", b.ContentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(b.Body)))
	w.WriteHeader(b.StatusCode)
	if _, err := w.Write(b.Body); err != nil {
		log.Println("unable to write image.", err)
	}

}
