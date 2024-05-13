package database

type BookDetail struct {
	Name             string `json:"name"`
	Author           string `json:"author"`
	Publication_Year int    `json:"publication-year"`
}
