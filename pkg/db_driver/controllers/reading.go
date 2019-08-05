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
	for iter.Scan(&reading.Timestamp, &reading.NodeId, &reading.WaterLevel, &reading.Photo) {
		readings = append(readings, reading)
	}
	return readings, iter.Close()
}

func InsertReading(reading models.Reading) (bool, error) {
	return db_driver.GetDriver().GetSession().
		Query(`INSERT INTO readings (timestamp, node_id, water_level, photo) VALUES (?, ?, ?, ?) IF NOT EXISTS`).
		ScanCAS()
}
func DeleteReading(timestamp time.Time, nodeId string) (bool, error) {
	return db_driver.GetDriver().GetSession().
		Query(`DELETE FROM readings WHERE timestamp = ? and node_id = ? IF EXISTS`, timestamp, nodeId).ScanCAS()
}
