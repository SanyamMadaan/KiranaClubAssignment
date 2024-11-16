package models

import (
	"errors" // Importing the standard errors package
	"time"
)

type Job struct {
	JobID     string     `json:"job_id"`
	Visits    []Visit    `json:"visits"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	ErrorLogs []JobError `json:"error_logs,omitempty"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_urls"`
	VisitTime string   `json:"visit_time"`
}

type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

// JobMap is a global map to store jobs (simulate in-memory storage)
var JobMap = map[string]*Job{}

// CreateJob initializes a new job and stores it in JobMap
func CreateJob(visits []Visit) (*Job, error) {
	if len(visits) == 0 {
		return nil, errors.New("no visits provided")
	}

	job := &Job{
		JobID:     generateJobID(),
		Visits:    visits,
		Status:    "ongoing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	JobMap[job.JobID] = job
	return job, nil
}

// GetJobStatus retrieves the status of a job by jobID
func GetJobStatus(jobID string) (*Job, error) {
	job, exists := JobMap[jobID]
	if !exists {
		return nil, errors.New("job not found")
	}
	return job, nil
}

// UpdateJobStatus updates the status of a job (ongoing, completed, or failed)
func UpdateJobStatus(jobID, status string, jobErrors []JobError) error {
	job, exists := JobMap[jobID]
	if !exists {
		return errors.New("job not found")
	}

	job.Status = status
	job.UpdatedAt = time.Now()

	if len(jobErrors) > 0 {
		job.ErrorLogs = jobErrors
	}
	return nil
}

// Helper function to generate unique Job IDs (simple implementation for now)
func generateJobID() string {
	return time.Now().Format("20060102150405") // e.g., "20241117123045"
}
