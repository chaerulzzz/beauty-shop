package dto

type BookDTO struct {
	ID     uint   `json:"id,string,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
