package event

import (
	"event_project/domain/event"
	db "event_project/infrastructure/database/event"
	"net/http"
	"strings"
	"time"

	ua "github.com/mileusna/useragent"
	"goyave.dev/goyave/v3"
	"goyave.dev/goyave/v3/validation"
)

var (
	EventValidation validation.RuleSet = validation.RuleSet{
		"timestamp": {"required", "integer", "min:0"},
		"type":      {"required", "string", "eventType"},
	}
)

func SaveEvent(response *goyave.Response, request *goyave.Request) {
	userAgent := ua.Parse(request.UserAgent())
	timestamp := time.Unix(int64(request.Integer("timestamp")), 0)
	dbHandler := db.EventDatabaseHander{}
	t, err := event.TypeFromString(request.String("type"))
	if err != nil {
		if internalError := response.Error(err); internalError != nil {
			response.JSON(http.StatusInternalServerError, nil)
			return
		}
	}

	event := event.EventEntity{
		Type:      t,
		Timestamp: timestamp.Unix(),
		UserAgent: request.UserAgent(),
		Os:        userAgent.OS,
		Ip:        parsePortFromRemoveAddr(request.Request().RemoteAddr),
	}

	if err := dbHandler.Save(&event); err != nil {
		response.Error(err)
		return
	}

	response.JSON(http.StatusCreated, nil)
}

func parsePortFromRemoveAddr(ipPort string) string {
	return strings.Split(ipPort, ":")[0]
}
