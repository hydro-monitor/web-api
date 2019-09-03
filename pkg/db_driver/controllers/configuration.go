package controllers

import (
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/models"
)

func GetNodeConfiguration(nodeId string) (models.Configuration, error) {
	var configuration models.Configuration
	var state string
	err := db_driver.GetDriver().GetSession().
		Query(`SELECT state FROM nodes WHERE node_id = ?`, nodeId).Scan(&state)
	if err != nil {
		return configuration, err
	}
	// TODO see why order of columns does not match model
	err = db_driver.GetDriver().GetSession().
		Query(`SELECT * FROM states WHERE node_id = ? and name = ?`, nodeId, state).
		Scan(&configuration.NodeId,
			&configuration.Name,
			&configuration.MsBetweenReadings,
			&configuration.NextState,
			&configuration.PhotosPerReading,
			&configuration.PreviousState,
			&configuration.WaterLevelLimitForGoingToNextState,
			&configuration.WaterLevelLimitForGoingToPreviousState)
	return configuration, err
}
