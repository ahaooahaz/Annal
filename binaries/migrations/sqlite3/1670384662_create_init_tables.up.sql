CREATE TABLE IF NOT EXISTS `todos` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `uuid` varchar(64) NOT NULL UNIQUE,
    `title` varchar(256) NOT NULL,
    `description` text NOT NULL,
    `plan` datetime(6) NOT NULL,
    `status` integer NOT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
);