package entity

import "github.com/LuanTenorio/learn-api/internal/database/sqlc"

type Subject struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	TotalTime int    `json:"total_time" db:"total_time"`
	Avarage   int    `json:"avarage"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UserId    int    `json:"-" db:"user_id"`
}

func M2E(m *sqlc.Subject) *Subject {
	return &Subject{
		Id:        int(m.ID),
		Name:      m.Name,
		TotalTime: int(m.TotalTime),
		Avarage:   int(m.Avarage),
		UserId:    int(m.UserID),
		CreatedAt: m.CreatedAt.Time.String(),
	}
}
