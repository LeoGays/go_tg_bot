package repository

import (
	"context"
	"go-tg-bot/internal/domain"
	"go-tg-bot/internal/repository/ent"
)

func (r Repository) CreateEvent(ctx context.Context, event *domain.Event) (*ent.Event, error) {
	eventEnt, err := r.Client.Event.
		Create().
		SetName(event.Name).
		SetDescription(event.Description).
		SetLocation(event.Location).
		SetDate(event.Date).
		Save(ctx)
	
	if err != nil {
		return nil, err
	}
	return eventEnt, nil
}
