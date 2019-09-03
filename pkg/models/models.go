package models

import "time"

type Configuration struct {
	NodeId                                 string
	Name                                   string
	PhotosPerReading                       int
	MsBetweenReadings                      int
	WaterLevelLimitForGoingToPreviousState float32
	WaterLevelLimitForGoingToNextState     float32
	PreviousState                          string
	NextState                              string
}

type Reading struct {
	NodeId     string    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	Photo      []byte
	WaterLevel float32 `json:"water_level"`
}

type Node struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	State       string `json:"state"`
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
