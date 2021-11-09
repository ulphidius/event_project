package event

import (
	"strings"
	"time"

	"goyave.dev/goyave/v3/database"
)

func init() {
	database.RegisterModel(&EventEntity{})
}

type EventEntity struct {
	ID        uint   `gorm:"primary_key;auto_increment;not_null"`
	Type      Type   `gorm:"size:255"`
	UserAgent string `gorm:"size:500"`
	Ip        string `gorm:"size:11"`
	Timestamp int64
	Os        string `gorm:"size:50"`
}

type OsStatistics struct {
	NumberOfLinux   uint `json:"number_of_linux"`
	NumberOfMac     uint `json:"number_of_mac"`
	NumberOfWindows uint `json:"number_of_windows"`
	NumberOfPhone   uint `json:"number_of_phone"`
	Other           uint `json:"other"`
}

type TypeStatistics struct {
	NumberOfImpression uint `json:"number_of_impression"`
	NumberOfClick      uint `json:"number_of_click"`
	NumberOfVisible    uint `json:"number_of_visible"`
}

func GetOsStatistics(events []EventEntity, start time.Time, end time.Time) OsStatistics {
	stat := OsStatistics{0, 0, 0, 0, 0}
	events = filterByTime(events, start, end)

	for _, event := range events {
		switch strings.ToLower(event.Os) {
		case "linux":
			stat.NumberOfLinux += 1
		case "windows":
			stat.NumberOfWindows += 1
		case "macos":
			stat.NumberOfMac += 1
		case "android":
			stat.NumberOfPhone += 1
		case "iphone":
			stat.NumberOfPhone += 1
		default:
			stat.Other += 1
		}
	}

	return stat
}

func GetTypeStatistics(events []EventEntity, start time.Time, end time.Time) TypeStatistics {
	stat := TypeStatistics{0, 0, 0}
	events = filterByTime(events, start, end)

	for _, event := range events {
		switch event.Type {
		case Impression:
			stat.NumberOfImpression += 1
		case Click:
			stat.NumberOfClick += 1
		case Visible:
			stat.NumberOfVisible += 1
		}
	}

	return stat
}

func filterByTime(events []EventEntity, start time.Time, end time.Time) []EventEntity {
	var filtered_event []EventEntity

	for _, event := range events {
		eventDate := time.Unix(event.Timestamp, 0)
		if (eventDate.After(start) || eventDate.Equal(start)) && (eventDate.Before(end) || eventDate.Equal(end)) {
			filtered_event = append(filtered_event, event)
		}
	}

	return filtered_event
}
