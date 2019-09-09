package controllers

import (
	"errors"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/models"
)

func UpdateManualReading(manualReading models.ManualReading) error {
	stmt, names := qb.Update("hydromonitor.manual_readings").Set("reading_required").Where(qb.Eq("node_id")).Existing().ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).BindStruct(manualReading)
	applied, err := q.ScanCAS()
	if !applied {
		return errors.New("Specified node id does not exist")
	}
	return err
}

func GetManualReading(nodeId string) (models.ManualReading, error) {
	var manualReading models.ManualReading
	stmt, names := qb.Select("hydromonitor.manual_readings").Where(qb.Eq("node_id")).ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).Bind(nodeId)
	err := q.GetRelease(&manualReading)
	return manualReading, err
}
