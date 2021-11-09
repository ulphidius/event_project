package event

import (
	"event_project/domain/event"

	"gorm.io/gorm"
	"goyave.dev/goyave/v3/database"
)

type EventDatabaseHander struct {
	db *gorm.DB
}

func (handler *EventDatabaseHander) Save(event *event.EventEntity) error {
	if handler.db == nil {
		handler.db = database.Conn()
	}

	err := handler.db.Create(&event).Error

	return err
}

func (handler *EventDatabaseHander) GetAll() ([]event.EventEntity, error) {
	var events []event.EventEntity

	if handler.db == nil {
		handler.db = database.Conn()
	}

	err := handler.db.Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}
