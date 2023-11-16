package api

import (
	"dz/Structs"
	"net/http"
)

type api struct {
	addr string
	r    *http.ServeMux
	s    *Structs.LocalStorage
}

func New(addr string) *api {
	return &api{addr: addr, r: new(http.ServeMux), s: Structs.New()}
}

func (api *api) FillEndpoints() {
	api.r.HandleFunc("/api/answer", api.AnswerHandle)
	api.r.HandleFunc("/api/question", api.QuestionHandle)
}
func (api *api) ListenAndServe() error {
	return http.ListenAndServe(api.addr, api.r)
}
