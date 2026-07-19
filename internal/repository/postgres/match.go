package postgres

import (
	"go-games-microservice/internal/domain"
	"time"

	"gorm.io/gorm"
)

type MatchDB struct {
	gorm.Model
	Status    domain.MatchStatus
	StartedAt time.Time
	EndedAt   *time.Time
	Events    []MatchEventDB `gorm:"foreignKey:MatchID"`
}

func (MatchDB) TableName() string {
	return "matches"
}

func (m *MatchDB) ToDomain() *domain.Match {
	return &domain.Match{
		ID:        int(m.ID),
		Status:    m.Status,
		StartedAt: m.StartedAt,
		EndedAt:   m.EndedAt,
	}
}

func MatchFromDomain(match domain.Match) *MatchDB {
	return &MatchDB{
		Model: gorm.Model{
			ID: uint(match.ID),
		},
		Status:    match.Status,
		StartedAt: match.StartedAt,
		EndedAt:   match.EndedAt,
	}
}

type MatchRepository struct {
	db *DB
}

func NewMatchRepository(db *DB) *MatchRepository {
	return &MatchRepository{
		db: db,
	}
}
