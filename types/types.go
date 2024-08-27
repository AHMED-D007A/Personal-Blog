package types

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
}
