package httphandler

import (
	"net/http"

	"github.com/hasemeneh/PoC-OnlineStore/helper/response"
	"github.com/julienschmidt/httprouter"
)

type HttpHandler interface {
	GET(path string, handler handlerFunc)
	POST(path string, handler handlerFunc)
	PUT(path string, handler handlerFunc)
	OPTIONS(path string, handler handlerFunc)
}

type httpHandler struct {
	httpRouter *httprouter.Router
	prefix     string
}

func New(prefix string, httpRouter *httprouter.Router) HttpHandler {
	return &httpHandler{
		httpRouter: httpRouter,
		prefix:     prefix,
	}
}

type handlerFunc func(r *http.Request) response.HttpResponse

func (h *httpHandler) GET(path string, handler handlerFunc) {
	h.register(http.MethodGet, path, handler)
}

func (h *httpHandler) POST(path string, handler handlerFunc) {
	h.register(http.MethodPost, path, handler)
}

func (h *httpHandler) PUT(path string, handler handlerFunc) {
	h.register(http.MethodPut, path, handler)
}

func (h *httpHandler) OPTIONS(path string, handler handlerFunc) {
	h.register(http.MethodOptions, path, handler)
}
func (h *httpHandler) register(method, path string, handler handlerFunc) {
	fullpath := h.prefix + path
	h.httpRouter.Handle(method, fullpath, transformator(handler))
}

func transformator(handler handlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		resp := handler(r)
		resp.WriteResponse(w)
	}
}
