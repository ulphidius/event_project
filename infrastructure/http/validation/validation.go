package validation

import (
	"event_project/domain/event"
	"strings"

	"goyave.dev/goyave/v3/validation"
)

// If none of the available validation rules satisfy your needs, you can implement custom validation rules.
// https://goyave.dev/guide/basics/validation.html#custom-rules
func init() {
	validation.AddRule(
		"eventType",
		&validation.RuleDefinition{
			Function:           validateEventType,
			RequiredParameters: 0,
		})
}

func validateEventType(field string, value interface{}, parameters []string, form map[string]interface{}) bool {
	str, ok := value.(string)
	str = strings.Title(strings.ToLower(str))

	if ok {
		return event.TypeExist(str)
	}

	return false
}
