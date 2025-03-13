package internal

type Course struct {
	Id          *int   `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
