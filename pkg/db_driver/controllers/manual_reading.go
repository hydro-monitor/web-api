package controllers

import (
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/models"
)

func UpdateManualReading(manualReading models.ManualReading) (bool, error) {
	return db_driver.GetDriver().GetSession().
		Query(`UPDATE manual_readings SET reading_required = ? WHERE node_id = ? IF EXISTS`,
			manualReading.ReadingRequired, manualReading.NodeId).
		ScanCAS()
}

func GetManualReading(nodeId string) (models.ManualReading, error) {
	var manualReading models.ManualReading
	err := db_driver.GetDriver().GetSession().Query(`SELECT * FROM manual_readings WHERE node_id = ?`, nodeId).
		Scan(&manualReading.NodeId, &manualReading.ReadingRequired)
	return manualReading, err
}
