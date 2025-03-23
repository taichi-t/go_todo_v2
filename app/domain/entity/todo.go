package entity

type TodoStatus string

const (
	TodoStatusNotStarted TodoStatus = "NOT_STARTED"
	TodoStatusInProgress TodoStatus = "IN_PROGRESS"
	TodoStatusCompleted  TodoStatus = "COMPLETED"
)

type Todo struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TodoStatus `json:"status"`
}
