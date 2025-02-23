package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (s *ApiServer) getAttributeDefs(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := s.attributeDefMgmt.GetAttributeDefs()
	if err != nil {
		slog.Error("GetAttributeDefs failed.", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("GetHealthCheck failed to encode response as json.", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
