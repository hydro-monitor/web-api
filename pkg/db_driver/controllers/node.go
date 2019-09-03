package controllers

import (
	"github.com/gocql/gocql"
	"hydro_monitor/web_api/pkg/db_driver"
	"hydro_monitor/web_api/pkg/models"
)

func GetAllNodes() ([]models.Node, error) {
	var node models.Node
	var nodes []models.Node
	iter := db_driver.GetDriver().GetSession().Query(`SELECT * FROM nodes`).Consistency(gocql.One).Iter()
	for iter.Scan(&node.Id, &node.Description, &node.State) {
		nodes = append(nodes, node)
	}
	err := iter.Close()
	return nodes, err
}

func GetNodeByID(id string) (models.Node, error) {
	var node models.Node
	err := db_driver.GetDriver().GetSession().Query(`SELECT * FROM nodes WHERE node_id = ?`, id).Consistency(gocql.One).Scan(&node.Id, &node.Description, &node.State)
	return node, err
}

func InsertNode(node models.Node) (bool, error) {
	return db_driver.GetDriver().GetSession().
		Query(`INSERT INTO nodes (node_id, description, state) VALUES (?, ?, ?) IF NOT EXISTS`,
			node.Id, node.Description, node.State).ScanCAS()
}

func DeleteNode(id string) (bool, error) {
	return db_driver.GetDriver().GetSession().Query(`DELETE FROM nodes WHERE node_id = ? IF EXISTS`, id).ScanCAS()
}
