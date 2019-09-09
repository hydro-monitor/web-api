package controllers

import (
	"errors"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/models"
)

func GetNodeConfiguration(nodeId string) (models.Configuration, error) {
	var configuration models.Configuration
	var state string
	stmt, names := qb.Select("hydromonitor.nodes").Columns("state").Where(qb.Eq("id")).ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).Bind(nodeId)
	if err := q.GetRelease(&state); err != nil {
		return configuration, err
	}

	stmt, names = qb.Select("hydromonitor.states").Columns("*").Where(qb.Eq("name")).ToCql()
	q = gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).Bind(state)
	err := q.GetRelease(&configuration)
	return configuration, err
}

func UpdateNodeConfiguration(configuration models.Configuration) error {
	stmt, names := qb.Update("hydromonitor.states").
		Set("ms_between_readings",
			"next_state",
			"photos_per_reading",
			"previous_state",
			"water_level_limit_for_next_state",
			"water_level_limit_for_previous_state").
		Where(qb.Eq("name")).Existing().ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).BindStruct(configuration)
	applied, err := q.ScanCAS()
	if !applied {
		return errors.New("Specified configuration does not exist")
	}
	return err
}
