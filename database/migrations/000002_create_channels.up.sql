CREATE TABLE `channels` (
    `id` VARCHAR(36) COLLATE utf8mb4_general_ci NOT NULL,
    `name` VARCHAR(255) COLLATE utf8mb4_general_ci NOT NULL,
    `owner_id` VARCHAR(36) COLLATE utf8mb4_general_ci NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `owner_id` (`owner_id`),
    CONSTRAINT `channels_owner_id_foreign` FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
