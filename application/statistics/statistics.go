package statistics

import (
	"event_project/domain/event"
	db "event_project/infrastructure/database/event"
	"net/http"
	"strconv"
	"strings"
	"time"

	"goyave.dev/goyave/v3"
)

func GetStatistics(response *goyave.Response, request *goyave.Request) {
	dbHandler := db.EventDatabaseHander{}
	if request.Data["type"] == nil || request.Data["min"] == nil || request.Data["max"] == nil {
		response.JSON(http.StatusUnprocessableEntity, "The arguments type, min and max are required")
		return
	}
	t := request.Data["type"].(string)
	if !typeIsValid(t) {
		response.JSON(http.StatusUnprocessableEntity, "For type argument only the values 'os' and 'type' are allowed")
		return
	}
	min, err := strconv.Atoi(request.Data["min"].(string))
	if err != nil {
		response.JSON(http.StatusUnprocessableEntity, "Min must be integer")
		return
	}
	max, err := strconv.Atoi(request.Data["max"].(string))
	if err != nil {
		response.JSON(http.StatusUnprocessableEntity, "Max must be integer")
		return
	}

	if min > max {
		response.JSON(http.StatusUnprocessableEntity, "Min must be lower or equal to max")
		return
	}
	if min < 0 || max < 0 {
		response.JSON(http.StatusUnprocessableEntity, "Min and max must be upper than 0")
		return
	}

	minTime := time.Unix(int64(min), 0)
	maxTime := time.Unix(int64(max), 0)
	t = strings.ToLower(t)
	events, err := dbHandler.GetAll()
	if err != nil {
		response.Error(err)
	}

	if t == "os" {
		statistics := event.GetOsStatistics(events, minTime, maxTime)
		response.JSON(http.StatusOK, statistics)
		return
	}
	if t == "type" {
		statistics := event.GetTypeStatistics(events, minTime, maxTime)
		response.JSON(http.StatusOK, statistics)
		return
	}

	response.JSON(http.StatusInternalServerError, nil)
}

func typeIsValid(t string) bool {
	tLowerCase := strings.ToLower(t)
	if tLowerCase == "os" || tLowerCase == "type" {
		return true
	}

	return false
}
