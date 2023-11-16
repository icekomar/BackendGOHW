package api

import (
	"dz/Structs"
	"encoding/json"
	"net/http"
)

func (api *api) QuestionHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_, err := w.Write([]byte(api.s.ToString(true)))
		if err != nil {
			return
		}

	case http.MethodPost:
		api.s.Lock()
		var m []Structs.Record
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			api.s.Unlock()
			return
		}
		api.s.AddFromJson(m)
		api.s.Unlock()
	}
}
