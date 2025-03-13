CREATE TABLE `users` (
    `id` VARCHAR(36) COLLATE utf8mb4_general_ci NOT NULL,
    `first_name` VARCHAR(255) COLLATE utf8mb4_general_ci NOT NULL,
    `last_name` VARCHAR(255) COLLATE utf8mb4_general_ci NOT NULL,
    `email` VARCHAR(255) COLLATE utf8mb4_general_ci NOT NULL,
    `password` VARCHAR(255) COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email_unique` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
