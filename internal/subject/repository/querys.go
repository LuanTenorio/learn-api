package repository

const (
	createSubjectQuery  = "INSERT INTO subjects (name, total_time, avarage, user_id) VALUES (:name, 0, 0, :user_id) RETURNING *;"
	findIdByNameQuery   = "SELECT id FROM subjects WHERE name=$1 and user_id=$2 LIMIT 1;"
	findPaginationQuery = "SELECT * FROM subjects LIMIT $1 OFFSET $2;"
	totalItemsQuery     = "SELECT COUNT(id) AS total FROM subjects WHERE user_id = $1;"
)
