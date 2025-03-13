package internal


type Response struct {
	Data Course `json:"data"`
}
type Course struct {
	Id          *int   `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
