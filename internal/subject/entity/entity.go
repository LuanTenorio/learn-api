package entity

type SubjectEntity struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	TotalTime int    `json:"total_time" db:"total_time"`
	Avarage   int    `json:"avarage"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
