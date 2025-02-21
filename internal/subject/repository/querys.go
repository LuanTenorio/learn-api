package repository

const (
	createSubjectQuery = "INSERT INTO subjects (name, total_time, avarage) VALUES (:name, 0, 0) RETURNING *;"
	findIdByName       = "SELECT id FROM subjects WHERE name=$1 LIMIT 1;"
)
