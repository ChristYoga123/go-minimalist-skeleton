package responses

// HealthResponse is the DTO for the health check endpoint.
// It standardizes the JSON structure of the response.
type HealthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}
