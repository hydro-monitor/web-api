package controllers

import (
	"github.com/gocql/gocql"
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/models"
	"time"
)

func GetAllReadingsFromNode(nodeId string) ([]models.Reading, error) {
	var reading models.Reading
	var readings []models.Reading
	iter := db_driver.GetDriver().GetSession().Query(`SELECT * FROM readings WHERE node_id = ?`, nodeId).
		Consistency(gocql.One).Iter()
	for iter.Scan(&reading.NodeId, &reading.Timestamp, &reading.Photo, &reading.WaterLevel) {
		readings = append(readings, reading)
	}
	return readings, iter.Close()
}

func InsertReading(reading models.Reading) (bool, error) {
	return db_driver.GetDriver().GetSession().
		Query(`INSERT INTO readings (node_id, reading_time, photo, water_level) VALUES (?, ?, ?, ?) IF NOT EXISTS`,
			reading.NodeId, reading.Timestamp, reading.Photo, reading.WaterLevel).
		ScanCAS()
}
func DeleteReading(timestamp time.Time, nodeId string) (bool, error) {
	return db_driver.GetDriver().GetSession().
		Query(`DELETE FROM readings WHERE node_id = ? and reading_time = ? IF EXISTS`, nodeId, timestamp).ScanCAS()
}
