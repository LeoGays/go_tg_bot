package domain

import (
	"github.com/google/uuid"
	"go-tg-bot/internal/repository/ent"
	"time"
)

type Event struct {
	ID          uuid.UUID
	Name        string
	Description string
	Location    string
	Date        time.Time
}

func ConvertEventEntToEventDomain(event *ent.Event) *Event {
	return &Event{
		ID:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		Date:        event.Date,
	}
}
