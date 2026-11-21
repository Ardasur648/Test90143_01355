package scripts

import (
	"encoding/json"
	"net/http"
)

type BridgePayload struct {
	Event string         `json:"event"`
	Data  map[string]any `json:"data"`
}

func RegisterSiteHandlers(mux *http.ServeMux, siteDir string) {
	mux.Handle("/", http.FileServer(http.Dir(siteDir)))
	mux.HandleFunc("/bridge", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var payload BridgePayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status": "ok",
			"event":  payload.Event,
		})
	})
}