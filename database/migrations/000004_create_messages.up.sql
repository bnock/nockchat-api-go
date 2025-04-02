CREATE TABLE `messages` (
    `id` VARCHAR(36) COLLATE utf8mb4_general_ci NOT NULL,
    `channel_id` VARCHAR(36) COLLATE utf8mb4_general_ci NOT NULL,
    `sender_id` VARCHAR(36) COLLATE utf8mb4_general_ci NOT NULL,
    `content` TEXT COLLATE utf8mb4_general_ci NOT NULL,
    `sent_at` TIMESTAMP NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    `deleted_at` TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `channel_id` (`channel_id`),
    KEY `sender_id` (`sender_id`),
    CONSTRAINT `messages_channel_id_foreign` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON DELETE CASCADE,
    CONSTRAINT `messages_sender_id_foreign` FOREIGN KEY (`sender_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
