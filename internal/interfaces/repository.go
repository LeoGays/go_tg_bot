package interfaces

import (
	"context"
	"go-tg-bot/internal/domain"
	"go-tg-bot/internal/repository/ent"
)

type RepositoryEvent interface {
	CreateEvent(ctx context.Context, event *domain.Event) (*ent.Event, error)
}

type Repository interface {
	RepositoryEvent
}
