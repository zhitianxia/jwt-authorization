package model

type ProblemModel struct {
	Status int `json:"status"`
	Title string `json:"title"`
	Detail string `json:"detail"`
}
