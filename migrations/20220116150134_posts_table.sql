-- +goose Up
-- +goose StatementBegin
CREATE TABLE `posts` (
 `id` bigint unsigned NOT NULL AUTO_INCREMENT,
 `user_id` bigint unsigned NOT NULL AUTO_INCREMENT,
 `title` varchar(512) NOT NULL DEFAULT '',
 `abstract` text NOT NULL,
 `article` mediumtext NOT NULL,
 `keywords` varchar(512) NOT NULL DEFAULT '',
 `category` tinyint NOT NULL,
 `slug` varchar(255) NOT NULL DEFAULT '',
 `allow_comments` tinyint(1) NOT NULL DEFAULT '1',
 `tags` varchar(512) NOT NULL DEFAULT '',
 `photo` mediumtext NOT NULL,
 PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE `posts`
-- +goose StatementEnd
