package event

import "errors"

type Type int

const (
	Impression Type = iota
	Click
	Visible
)

func TypeExist(t string) bool {
	switch t {
	case "Impression":
		return true
	case "Click":
		return true
	case "Visible":
		return true
	}

	return false
}

func TypeFromString(t string) (Type, error) {
	switch t {
	case "Impression":
		return 0, nil
	case "Click":
		return 1, nil
	case "Visible":
		return 2, nil
	}

	return -1, errors.New("This event type doesn't exist")
}

func (t Type) String() string {
	switch t {
	case Impression:
		return "Impression"
	case Click:
		return "Click"
	case Visible:
		return "Visible"
	}

	return "Unknown"
}
