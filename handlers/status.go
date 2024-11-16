package handlers

import (
	"encoding/json"
	"net/http"
)

func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	if jobID == "" {
		http.Error(w, "Missing jobid", http.StatusBadRequest)
		return
	}

	status, exists := Jobs[jobID]
	if !exists {
		http.Error(w, "Job ID not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"job_id": jobID,
		"status": status,
	})
}
