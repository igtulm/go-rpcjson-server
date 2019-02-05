package main

type User struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	CreatedAt int64  `json:"created_at"`
}
