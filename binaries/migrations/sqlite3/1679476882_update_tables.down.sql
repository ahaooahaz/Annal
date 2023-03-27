BEGIN;
CREATE TABLE `tmp_todos` AS SELECT `id`, `uuid`, `idx`, `title`, `description`, `plan`, `status`, `created_at`, `updated_at` FROM `todos`;
DROP TABLE IF EXISTS `todos`;
ALTER TABLE `tmp_todos` RENAME TO `todos`;
COMMIT;