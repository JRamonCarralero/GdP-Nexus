package types

type ProjectRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Owner       string   `json:"owner" binding:"required"`
	Members     []string `json:"members"`
}

type ProjectUpdateRequest struct {
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Owner       *string   `json:"owner,omitempty"`
	Members     *[]string `json:"members,omitempty"`
}
