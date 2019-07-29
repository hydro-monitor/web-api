package controllers

import (
	"github.com/gocql/gocql"
	"hydro_monitor/web_api/db_driver"
	"hydro_monitor/web_api/models"
)

func GetAllNodes() ([]models.Node, error) {
	var node models.Node
	var nodes []models.Node
	iter := db_driver.GetDriver().GetSession().Query("SELECT * FROM nodes").Consistency(gocql.One).Iter()
	for iter.Scan(&node.Id, &node.Description) {
		nodes = append(nodes, node)
	}
	err := iter.Close()
	return nodes, err
}

func GetNodeByID(id string) (models.Node, error) {
	var node models.Node
	err := db_driver.GetDriver().GetSession().Query("SELECT * FROM nodes WHERE ID = ?", id).Consistency(gocql.One).Scan(&node.Id, &node.Description)
	return node, err
}
