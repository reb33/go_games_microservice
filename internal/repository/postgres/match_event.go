package postgres

import (
	"go-games-microservice/internal/domain"
	"time"

	"gorm.io/gorm"
)

type MatchEventDB struct {
	gorm.Model
	Type        domain.EventType
	PlayerID    uint
	TargetID    *uint
	TimeOfEvent time.Time
	MatchID     uint
}

func (MatchEventDB) TableName() string {
	return "matches_events"
}

func (m *MatchEventDB) ToDomain() *domain.MatchEvent {
	var targetID *int
	if m.TargetID != nil {
		v := int(*m.TargetID)
		targetID = &v
	}
	return &domain.MatchEvent{
		ID:          int(m.ID),
		Type: m.Type,
		PlayerID: int(m.PlayerID),
		TargetID: targetID,
		TimeOfEvent: m.TimeOfEvent,
		MatchID: int(m.MatchID),
	}
}

func MatchEventFromDomain(matchEvent *domain.MatchEvent) *MatchEventDB {
	var targetID *uint
	if matchEvent.TargetID != nil {
		v := uint(*matchEvent.TargetID)
		targetID = &v
	}
	return &MatchEventDB{
		Model: gorm.Model{
			ID: uint(matchEvent.ID),
		},
		Type: matchEvent.Type,
		PlayerID: uint(matchEvent.PlayerID),
		TargetID: targetID,
		TimeOfEvent: matchEvent.TimeOfEvent,
		MatchID: uint(matchEvent.MatchID),
	}
}

type MatchEventRepository struct {
	db *DB
}

func NewMatchEventRepository(db *DB) *MatchEventRepository {
	return &MatchEventRepository{
		db: db,
	}
}
