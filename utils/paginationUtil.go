package utils

// PaginationRequest represents the pagination parameters from request
type PaginationRequest struct {
	Page  int `query:"page" validate:"min=1"`
	Limit int `query:"limit" validate:"min=1,max=100"`
}

// PaginationResponse represents the pagination metadata in response
type PaginationResponse struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}

// NewPaginationRequest creates a new pagination request with default values
func NewPaginationRequest() *PaginationRequest {
	return &PaginationRequest{
		Page:  1,
		Limit: 10,
	}
}

// GetOffset calculates the offset for database query
func (p *PaginationRequest) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

// NewPaginationResponse creates a new pagination response
func NewPaginationResponse(page, limit int, totalItems int64) *PaginationResponse {
	totalPages := int(totalItems) / limit
	if int(totalItems)%limit > 0 {
		totalPages++
	}

	return &PaginationResponse{
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
	}
} 