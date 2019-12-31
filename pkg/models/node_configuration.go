package models

type NodeConfiguration struct {
	NodeId string            `json:"nodeId"`
	States map[string]*State `json:"states"`
}
