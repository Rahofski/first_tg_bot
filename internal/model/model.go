package model

import "time"

//статья как элемент ленты (RSS)
type Item struct {
	Title      string
	Categories []string
	Link       string
	Date       time.Time
	Summary    string
	SourceName string
}

type Source struct {
	ID        int
	Name      string
	FeedURL   string
	CreatedAt time.Time
}


//статья как элемент БД
type Article struct {
	ID          int
	SourceID    int
	Title       string
	Link        string
	Summary     string
	PublishedAt time.Time
	PostedAt    time.Time
	CreatedAt   time.Time
}
