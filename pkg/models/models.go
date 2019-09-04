package models

import "time"

type Configuration struct {
	NodeId                                 string  `json:"nodeId"`
	Name                                   string  `json:"name"`
	PhotosPerReading                       int     `json:"photosPerReading"`
	MsBetweenReadings                      int     `json:"msBetweenReadings"`
	WaterLevelLimitForGoingToPreviousState float32 `json:"waterLevelLimitForGoingToPreviousState"`
	WaterLevelLimitForGoingToNextState     float32 `json:"waterLevelLimitForGointToNextState"`
	PreviousState                          string  `json:"previousState"`
	NextState                              string  `json:"nextState"`
}

type Reading struct {
	NodeId     string    `json:"nodeId"`
	Timestamp  time.Time `json:"timestamp"`
	Photo      []byte    `json:"photo"`
	WaterLevel float32   `json:"waterLevel"`
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
	NodeId          string `json:"nodeId"`
	ReadingRequired bool   `json:"readingRequired"`
}
