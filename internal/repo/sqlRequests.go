package repo

const (
	CreateTask = `INSERT INTO tasks (user_id, title, description) VALUES ($1, $2, $3) RETURNING id`
	GetTask    = `SELECT id, user_id, title, description, status, created_at FROM tasks WHERE id = $1`
	UpdateTask = `
	UPDATE tasks
	SET 
    title = COALESCE(NULLIF($2, ''), title), 
    description = COALESCE(NULLIF($3, ''), description), 
    status = COALESCE(NULLIF($4, ''), status),
    user_id = COALESCE(NULLIF($5, 0), user_id)
	WHERE id = $1
	RETURNING id;`
	DeleteTask = `DELETE FROM tasks WHERE id = $1 RETURNING id;`
)
