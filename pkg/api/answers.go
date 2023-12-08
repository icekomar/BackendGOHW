package api

import (
	"dz/pkg/models"
	"encoding/json"
	"net/http"
)

func (api *api) AnswerHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_, err := w.Write([]byte(api.s.ToString(false)))
		if err != nil {
			return
		}

	case http.MethodPost:
		var m []models.Record
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			return
		}
		api.s.AddFromJson(m)
	}
}
