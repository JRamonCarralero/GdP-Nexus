package types

type ProjectRequest struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Owner       string   `json:"owner" binding:"required"`
	Members     []string `json:"members"`
}
