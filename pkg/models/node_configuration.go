package models

import "hydro_monitor/web_api/pkg/models/api_models"

type NodeConfiguration struct {
	NodeId string                       `json:"nodeId"`
	States map[string]*api_models.State `json:"states"`
}
