-- +goose Up
-- +goose StatementBegin
CREATE TABLE `categories` (
  `id` tinyint unsigned ,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  `categories`
-- +goose StatementEnd
