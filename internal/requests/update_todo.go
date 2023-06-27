package requests

type UpdateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Closed      bool   `json:"closed" binding:"required"`
}
