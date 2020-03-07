package db_models

import (
	"github.com/scylladb/gocqlx/qb"
	"hydro_monitor/web_api/pkg/models/api_models"
)

var statesColumns = []string{"name", "photos_per_reading", "reading_interval", "lower_limit", "upper_limit",
	"previous_state", "next_state"}

type StatesDTO struct {
	nodeId string
	States []*StateDTO
}

func (s *StatesDTO) GetColumns() []string {
	return statesColumns
}

func (s *StatesDTO) GetBindMap() qb.M {
	return qb.M{"node_id": s.nodeId}
}

func (s *StatesDTO) GetArgs() interface{} {
	// We pass the pointer of the array to avoid the "must pass a pointer, not a value, to StructScan destination" error
	return &s.States
}

func (s *StatesDTO) ConvertToAPIStates() []*api_models.State {
	var states []*api_models.State
	for _, s := range s.States {
		states = append(states, s.ConvertToAPIState())
	}
	return states
}

func NewStatesDTO(nodeId string) *StatesDTO {
	states := make([]*StateDTO, 0)
	return &StatesDTO{
		nodeId: nodeId,
		States: states,
	}
}
