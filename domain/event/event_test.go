package event

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func EventEntitiesGenerator(number uint) []EventEntity {
	var events []EventEntity

	for i := 0; i < int(number); i++ {
		events = append(events, EventEntityGenerator())
	}

	return events
}

func EventEntityGenerator() EventEntity {
	listOs := []string{"Windows", "iPhone", "Android", "Unknown"}
	event := EventEntity{}

	rand.Seed(time.Now().Unix())
	event.Os = listOs[rand.Intn(len(listOs)-0+1)]

	event.Type = Type(rand.Intn(3 - 0 + 1))
	event.Ip = faker.IPv4()

	faker.SetGenerateUniqueValues(true)
	id, err := faker.RandomInt(0, 1000)
	if err != nil {
		panic("Cannot create ID")
	}

	event.ID = uint((id[0]))
	faker.SetGenerateUniqueValues(false)

	timestamp, err := time.Parse("2006-01-02 15:04:05", faker.Timestamp())
	if err != nil {
		panic("Cannot create Timestamp")
	}
	event.Timestamp = timestamp.Unix()

	return event
}

func SetRangeOfType(events []EventEntity, t Type, start uint, end uint) []EventEntity {
	for i := start; i < end; i++ {
		events[i].Type = t
	}

	return events
}

func SetRangeOfOs(events []EventEntity, os string, start uint, end uint) []EventEntity {
	for i := start; i < end; i++ {
		events[i].Os = os
	}

	return events
}

func SetRangeOfDate(events []EventEntity, date int64, start uint, end uint) []EventEntity {
	for i := start; i < end; i++ {
		events[i].Timestamp = date
	}

	return events
}

func filterData() []EventEntity {
	return []EventEntity{
		{
			ID:        0,
			Type:      Impression,
			Ip:        "127.0.0.1",
			Timestamp: time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			Os:        "Windows",
		},
		{
			ID:        1,
			Type:      Impression,
			Ip:        "127.0.0.1",
			Timestamp: time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			Os:        "Windows",
		},
		{
			ID:        2,
			Type:      Impression,
			Ip:        "127.0.0.1",
			Timestamp: time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			Os:        "Windows",
		},
		{
			ID:        3,
			Type:      Impression,
			Ip:        "127.0.0.1",
			Timestamp: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			Os:        "Windows",
		},
		{
			ID:        4,
			Type:      Impression,
			Ip:        "127.0.0.1",
			Timestamp: time.Date(2013, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			Os:        "Windows",
		},
		{
			ID:        5,
			Type:      Impression,
			Ip:        "127.0.0.1",
			Timestamp: time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			Os:        "Windows",
		},
		{
			ID:        6,
			Type:      Impression,
			Ip:        "127.0.0.1",
			Timestamp: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			Os:        "Windows",
		},
		{
			ID:        7,
			Type:      Impression,
			Ip:        "127.0.0.1",
			Timestamp: time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
			Os:        "Windows",
		},
	}
}
func TestFilterByTime(t *testing.T) {
	start := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)
	events := filterData()

	assert.Equal(t, events[1:7], filterByTime(events, start, end))
}

func TestFilterByTimeNoResult(t *testing.T) {
	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	events := filterData()

	assert.Nil(t, filterByTime(events, start, end))
}

func TestGetOsStatistics(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	events := SetRangeOfOs(
		SetRangeOfOs(filterData(), "Linux", 0, 8),
		"Windows",
		5,
		8,
	)
	expected := OsStatistics{5, 0, 3, 0, 0}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetOsStatisticsLinux(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	events := SetRangeOfOs(filterData(), "Linux", 0, 8)
	expected := OsStatistics{8, 0, 0, 0, 0}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetOsStatisticsWindows(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	events := filterData()
	expected := OsStatistics{0, 0, 8, 0, 0}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetOsStatisticsMacOs(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	events := SetRangeOfOs(filterData(), "MacOs", 0, 8)
	expected := OsStatistics{0, 8, 0, 0, 0}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetOsStatisticsAndroid(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	events := SetRangeOfOs(filterData(), "Android", 0, 8)
	expected := OsStatistics{0, 0, 0, 8, 0}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetOsStatisticsIPhone(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	events := SetRangeOfOs(filterData(), "iPhone", 0, 8)
	expected := OsStatistics{0, 0, 0, 8, 0}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetOsStatisticsOther(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	events := SetRangeOfOs(filterData(), "Sample", 0, 8)
	expected := OsStatistics{0, 0, 0, 0, 8}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetOsStatisticsEmptyEvents(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	var events []EventEntity = nil
	expected := OsStatistics{0, 0, 0, 0, 0}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetOsStatisticsOutOffDateRange(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2005, 1, 1, 0, 0, 0, 0, time.UTC)
	events := filterData()
	expected := OsStatistics{0, 0, 0, 0, 0}

	assert.Equal(t, expected, GetOsStatistics(events, start, end))
}

func TestGetTypeStatistics(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	events := SetRangeOfType(
		SetRangeOfType(filterData(), Click, 5, 7),
		Visible,
		7,
		8,
	)
	expected := TypeStatistics{5, 2, 1}

	assert.Equal(t, expected, GetTypeStatistics(events, start, end))
}
