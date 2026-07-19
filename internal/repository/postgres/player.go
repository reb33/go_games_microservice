package postgres

import (
	"go-games-microservice/internal/domain"
	"time"

	"gorm.io/gorm"
)

type PlayerDB struct {
	gorm.Model
	Nickname    string `gorm:"uniqueIndex;not null"`
	Name        string
	Lastname    string
	DateOfBirth time.Time
	Matches     []MatchDB `gorm:"many2many:player_matches;"`
}

func (PlayerDB) TableName() string {
	return "players"
}

func (p *PlayerDB) ToDomain() *domain.Player {
	return &domain.Player{
		ID:          int(p.ID),
		Nickname:    p.Nickname,
		Name:        p.Name,
		Lastname:    p.Lastname,
		DateOfBirth: p.DateOfBirth,
	}
}

func PlayerFromDomain(player *domain.Player) *PlayerDB {
	return &PlayerDB{
		Model: gorm.Model{
			ID: uint(player.ID),
		},
		Nickname:    player.Nickname,
		Name:        player.Name,
		Lastname:    player.Lastname,
		DateOfBirth: player.DateOfBirth,
	}
}

type PlayerRepository struct {
	db *DB
}

func NewPlayerRepository(db *DB) *PlayerRepository {
	return &PlayerRepository{
		db: db,
	}
}
