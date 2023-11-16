package api

import (
	"dz/Structs"
	"encoding/json"
	"fmt"
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
		api.s.Lock()
		var m []Structs.Record
		err := json.NewDecoder(r.Body).Decode(&m)
		if err != nil {
			return
		}
		fmt.Println(m)
		api.s.AddFromJson(m)
		api.s.Unlock()
	}
}
