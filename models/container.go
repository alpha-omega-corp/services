package models

type Container struct {
	ID    int64  `json:"id" bun:"id,pk"`
	Name  string `json:"name" bun:"name"`
	Image string `json:"image" bun:"image"`
}
