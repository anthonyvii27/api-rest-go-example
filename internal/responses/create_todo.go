package responses

import "time"

type CreateTodoResponse struct {
	Data struct {
		Id          string     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Closed      bool       `json:"closed"`
		CreatedAt   time.Time  `json:"createdAt"`
		UpdatedAt   *time.Time `json:"updatedAt"`
	} `json:"data"`
}
