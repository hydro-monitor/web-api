package services

import (
	"fmt"
	"sort"

	"hydro_monitor/web_api/pkg/models/api_models"
)

const defaultStateName = "default"

func ConfigToString(configuration map[string]*api_models.StateDTO) string {
	configStr := "[ "
	for stateName, state := range configuration {
		configStr += fmt.Sprintf("%s: %s ", stateName, stateToString(state))
	}
	configStr += "]"
	return configStr
}

func stateToString(state *api_models.StateDTO) string {
	return fmt.Sprintf("{interval: %d, lowerLimit: %f, upperLimit: %f, picturesNum: %d, prev: %s, next: %s}",
		state.Interval,
		state.LowerLimit,
		state.UpperLimit,
		state.PicturesNum,
		state.Prev,
		state.Next)
}

type ConfigurationValidator interface {
	ConfigurationIsValid(configuration map[string]*api_models.StateDTO) bool
}

type configurationValidatorImpl struct {}

func NewConfigurationValidator() ConfigurationValidator {
	return &configurationValidatorImpl{}
}

func (validator *configurationValidatorImpl) ConfigurationIsValid(configuration map[string]*api_models.StateDTO) bool {
	return defaultStateIsValid(configuration) && customStatesAreValid(configuration) && stateLimitsAreValid(configuration)
}

func defaultStateIsValid(configuration map[string]*api_models.StateDTO) bool {
	// Skip verification if there's no default state
	if _, present := configuration[defaultStateName]; !present {
		return true
	}
	defaultState := configuration[defaultStateName]
	// Default state can't have limits
	return defaultState.LowerLimit == 0 &&
		defaultState.UpperLimit == 0
}

func customStatesAreValid(configuration map[string]*api_models.StateDTO) bool {
	for stateName, state := range configuration {
		// Do not check default state
		if stateName == defaultStateName {
			continue
		}
		// Lower and upper limits can't be the same
		if state.LowerLimit == state.UpperLimit {
			fmt.Println("name:", stateName, "has matching lower and upper limit")
			return false
		}
	}
	return true
}

func stateLimitsAreValid(configuration map[string]*api_models.StateDTO) bool {
	var intervals [][]float64
	// Make intervals list
	for stateName, state := range configuration {
		// Do not check default state
		if stateName == defaultStateName {
			continue
		}
		fmt.Println("name:", stateName, "=>", "state:", state)
		interval := []float64{
			state.LowerLimit,
			state.UpperLimit,
		}
		intervals = append(intervals, interval)
	}
	// Sort intervals in increasing order of start time
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// In the sorted array, if start time of an interval
	// is less than end of previous interval, then there
	// is an overlap
	for i := 1; i < len(intervals); i++ {
		if intervals[i - 1][1] > intervals[i][0] {
			fmt.Println("overlap between intervals: ", intervals[i], intervals[i - 1])
			return false
		}
	}
	// If we reach here, then no overlap
	return true
}