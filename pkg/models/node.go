package models

type Node struct {
	Id            string `json:"id"`
	Description   string `json:"description"`
	Configuration string `json:"configuration"`
	State         string `json:"state"`
}
