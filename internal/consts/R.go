package consts

import "time"

type R struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []DataR `json:"data"`
}
type Data struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Extra string `json:"extra"`
	Link  string `json:"link"`
}
type DataR struct {
	ID         int       `json:"id"`
	Sort       int       `json:"sort"`
	Name       string    `json:"name"`
	SourceKey  string    `json:"source_key"`
	IconColor  string    `json:"icon_color"`
	Data       []Data    `json:"data"`
	CreateTime time.Time `json:"create_time"`
}
