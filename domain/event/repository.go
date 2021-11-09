package event

type EventRepository interface {
	Save(event *EventEntity) error
	GetAll() ([]EventEntity, error)
}
