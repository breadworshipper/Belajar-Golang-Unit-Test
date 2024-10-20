package entity

type Post struct {
	ID      int    `json:"id" gorm:"autoIncrement"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
