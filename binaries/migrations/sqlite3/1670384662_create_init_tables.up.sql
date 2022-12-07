CREATE TABLE IF NOT EXISTS `todos` (
    `id` integer PRIMARY KEY AUTOINCREMENT,
    `uuid` varchar(64) NOT NULL UNIQUE,
    `title` varchar(256) NOT NULL,
    `description` text NOT NULL,
    `plan` integer NOT NULL,
    `status` integer NOT NULL,
    `created_at` integer NOT NULL,
    `updated_at` integer NOT NULL
);