package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dxps/tmc-pwa/internal/shared/model"
)

func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := model.Health{State: "ok"}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		slog.Error("GetHealthCheck failed to encode response as json.", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
