package model

type Todo struct {
	Id          string `json:"id"`
	CreatedAt   string `json:"createdat"`
	Title       string `json:"title"`
	Discription string `json:"discription"`
}
