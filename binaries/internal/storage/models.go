package storage

type TodoTask struct {
	ID          int    `db:"id"`
	UUID        string `db:"uuid"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Plan        int64  `db:"plan"`
	Status      int    `db:"status"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
}
