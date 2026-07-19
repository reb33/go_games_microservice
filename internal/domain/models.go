package domain

import "time"

type MatchStatus string

const (
	MatchStatusNotStarted MatchStatus = "NOT_STARTED"
	MatchStatusInProgress MatchStatus = "IN_PROGRESS"
	MatchStatusFinished   MatchStatus = "FINISHED"
)

type EventType string

const (
	EventTypeKill   EventType = "KILL"
	EventTypeDeath  EventType = "DEATH"
	EventTypeAssist EventType = "ASSIST"
	EventTypeJoin   EventType = "JOIN"
	EventTypeLeave  EventType = "LEAVE"
)

type Player struct {
	ID          int
	Nickname    string
	Name        string
	Lastname    string
	DateOfBirth time.Time
}

type Match struct {
	ID        int
	Status    MatchStatus
	StartedAt time.Time
	EndedAt   *time.Time
}

type MatchEvent struct {
	ID          int
	MatchID     int
	Type        EventType
	PlayerID    int
	TargetID    *int
	TimeOfEvent time.Time
}
