-- +goose Up
-- +goose StatementBegin
CREATE TABLE `users` (
     `id` bigint NOT NULL AUTO_INCREMENT,
     `login` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
     `password_hash` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
     `password_reset_token` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci DEFAULT NULL,
     `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
     `role` smallint NOT NULL DEFAULT '1',
     `status` smallint NOT NULL DEFAULT '1',
     `created_at` datetime DEFAULT NULL,
     `updated_at` datetime DEFAULT NULL,
     UNIQUE KEY login (`login`),
     UNIQUE KEY email (`email`),
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users
-- +goose StatementEnd
