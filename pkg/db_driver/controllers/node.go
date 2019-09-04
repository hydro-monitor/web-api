package controllers

import (
	"errors"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/models"
)

func GetAllNodes() ([]models.Node, error) {
	var nodes []models.Node
	stmt, names := qb.Select("hydromonitor.nodes").ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names)
	err := q.SelectRelease(&nodes)
	return nodes, err
}

func GetNodeByID(id string) (models.Node, error) {
	var node models.Node
	stmt, names := qb.Select("hydromonitor.nodes").Where(qb.Eq("id")).ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).BindMap(qb.M{
		"id": id,
	})
	err := q.GetRelease(&node)
	return node, err
}

func InsertNode(node models.Node) error {
	stmt, names := qb.Insert("hydromonitor.nodes").Columns("id", "description", "state").Unique().ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).BindStruct(node)
	applied, err := q.ScanCAS()
	if !applied {
		return errors.New("Specified node id already exists")
	}
	return err
}

func DeleteNode(id string) error {
	stmt, names := qb.Delete("hydromonitor.nodes").Where(qb.Eq("id")).Existing().ToCql()
	q := gocqlx.Query(db_driver.GetDriver().GetSession().Query(stmt), names).Bind(id)
	applied, err := q.ScanCAS()
	if !applied {
		return errors.New("Specified node does not exist")
	}
	return err
}
