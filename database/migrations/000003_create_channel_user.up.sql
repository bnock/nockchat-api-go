CREATE TABLE `channel_user` (
    `channel_id` VARCHAR(36) COLLATE utf8mb4_general_ci NOT NULL,
    `user_id` VARCHAR(36) COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (`channel_id`,`user_id`),
    KEY `user_id` (`user_id`),
    CONSTRAINT `channel_user_channel_id_foreign` FOREIGN KEY (`channel_id`) REFERENCES `channels` (`id`) ON DELETE CASCADE,
CONSTRAINT `channel_user_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
