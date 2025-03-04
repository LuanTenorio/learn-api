-- name: CreateSubject :one
INSERT INTO subjects (name, total_time, avarage, user_id) VALUES ($1, 0, 0, $2) RETURNING *;

-- name: FindSubjectByIdAndName :one  
SELECT id FROM subjects WHERE name=$1 and user_id=$2 LIMIT 1;

-- name: ListSubjects :many
SELECT * FROM subjects WHERE user_id = $1 LIMIT $2 OFFSET $3;

-- name: TotalSubjectsByUser :one    
SELECT COUNT(id) AS total FROM subjects WHERE user_id = $1 LIMIT 1;