package api

import (
	"dz/pkg/storage"
	"net/http"
)

type api struct {
	addr string
	r    *http.ServeMux
	s    *storage.LocalStorage
}

func New(addr string) *api {
	return &api{addr: addr, r: new(http.ServeMux), s: storage.New()}
}

func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/answer", api.AnswerHandle)
	api.r.HandleFunc("/api/question", api.QuestionHandle)
}
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.r)
}
