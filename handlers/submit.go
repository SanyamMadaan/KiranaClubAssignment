package handlers

import (
	"encoding/json"
	"net/http"
	"store-image-processor/models"
)

type JobRequest struct {
	Count  int `json:"count"`
	Visits []struct {
		StoreID   string   `json:"store_id"`
		ImageURLs []string `json:"image_url"`
		VisitTime string   `json:"visit_time"`
	} `json:"visits"`
}

var Jobs = map[string]string{} // Track job statuses

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var req JobRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate Store IDs dynamically
	failedStores := []string{}
	for _, visit := range req.Visits {
		_, err := models.FetchStore(visit.StoreID)
		if err != nil {
			failedStores = append(failedStores, visit.StoreID)
		}
	}

	// If any store ID fails, return error
	if len(failedStores) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "failed",
			"error":  failedStores,
		})
		return
	}

	// Simulate Job Creation
	jobID := "job_" + fmt.Sprint(len(Jobs)+1)
	Jobs[jobID] = "ongoing"

	// Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"job_id": jobID,
		"status": "Job submitted successfully!",
	})
}
