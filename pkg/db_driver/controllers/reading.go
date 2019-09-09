package controllers

import (
	"errors"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/models"
	"time"
)

func GetAllReadingsFromNode(nodeId string) ([]models.Reading, error) {
	var readings []models.Reading
	stmt, names := qb.Select("hydromonitor.readings").ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names)
	err := q.SelectRelease(&readings)
	return readings, err
}

func InsertReading(reading models.Reading) error {
	stmt, names := qb.Insert("hydromonitor.readings").Columns("node_id", "reading_time", "photo", "water_level").Unique().ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).BindStruct(reading)
	applied, err := q.ScanCAS()
	if !applied {
		return errors.New("Specified reading already exists")
	}
	return err
}
func DeleteReading(timestamp time.Time, nodeId string) error {
	stmt, names := qb.Delete("hydromonitor.readings").Where(qb.Eq("node_id"), qb.Eq("reading_time")).Existing().ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).Bind(nodeId, timestamp)
	applied, err := q.ScanCAS()
	if !applied {
		return errors.New("Specified reading does not exist")
	}
	return err
}
