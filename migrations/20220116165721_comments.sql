-- +goose Up
-- +goose StatementBegin
CREATE TABLE `comments` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `post_id` bigint unsigned ,
    `comment_id` bigint unsigned ,
    `user_id` bigint unsigned ,
    `message` varchar(511) NOT NULL,
    `disable` tinyint(1) NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP `comments`
-- +goose StatementEnd
