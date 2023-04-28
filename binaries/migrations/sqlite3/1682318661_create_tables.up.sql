CREATE TABLE IF NOT EXISTS `words` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `word` varchar(32) NOT NULL UNIQUE,
    `pronunciation` varchar(64) NOT NULL,
    `cn` text NOT NULL,
    `status` integer NOT NULL DEFAULT 0,
    `meta` json NOT NULL,
    `created_at` integer NOT NULL,
    `updated_at` integer NOT NULL
);