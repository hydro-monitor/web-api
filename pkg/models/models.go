package models

import "time"

type Configuration struct {
	nodeId                                 string
	name                                   string
	photosPerReading                       int
	msBetweenReadings                      int
	waterLevelLimitForGoingToPreviousState int
	waterLevelLimitForGoingToNextState     int
	previousState                          string
	nextState                              string
}

type Reading struct {
	Timestamp  time.Time
	NodeId     string
	WaterLevel float64
	Photo      []byte
}

type Node struct {
	Id          string `json:"id" form:"id" query:"id"`
	Description string `json:"description" form:"description" query:"description"`
}

type User struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Admin    bool   `json:"admin" form:"admin" query:"admin"`
}

type ManualReading struct {
	NodeId          string
	ReadingRequired bool
}
