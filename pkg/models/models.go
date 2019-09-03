package models

import "time"

type Configuration struct {
	NodeId                                 string `json:"id"`
	Name                                   string `json:"name"`
	PhotosPerReading                       int    ``
	MsBetweenReadings                      int
	WaterLevelLimitForGoingToPreviousState float32
	WaterLevelLimitForGoingToNextState     float32
	PreviousState                          string
	NextState                              string
}

type Reading struct {
	NodeId     string    `json:"id"`
	Timestamp  time.Time `json:"timestamp"`
	Photo      []byte    `json:"photo"`
	WaterLevel float32   `json:"water_level"`
}

type Node struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	State       string `json:"state"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Admin    bool   `json:"admin"`
}

type ManualReading struct {
	NodeId          string `json:"id"`
	ReadingRequired bool   `json:"reading_required"`
}
