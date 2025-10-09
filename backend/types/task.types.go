package types

type TaskRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Project     string `json:"project" binding:"required"`
	Number      *int   `json:"number,omitempty"`
	Assignee    string `json:"assignee" binding:"required"`
	CreatedBy   string `json:"created_by" binding:"required"`
}

type TaskUpdateRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Assignee    *string `json:"assignee,omitempty"`
}
