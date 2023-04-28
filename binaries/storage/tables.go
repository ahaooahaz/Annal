package storage

var (
	todosTable = `todos`
	wordsTable = `words`
)

var (
	TodosCols = []string{`id`, `uuid`, `idx`, `title`, `description`, `plan`, `status`, `created_at`, `updated_at`, `notify_job_id`}
	wordsCols = []string{`id`, `word`, `pronunciation`, `cn`, `status`, `meta`, `created_at`, `updated_at`}
)
