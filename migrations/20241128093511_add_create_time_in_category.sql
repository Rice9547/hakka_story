-- +goose Up

ALTER TABLE story_to_category
    DROP PRIMARY KEY,
    ADD COLUMN id INT AUTO_INCREMENT PRIMARY KEY FIRST,
    ADD UNIQUE (story_id, category_id),
    ADD COLUMN created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- +goose Down

ALTER TABLE story_to_category
    DROP PRIMARY KEY,
    DROP COLUMN id,
    DROP COLUMN created_at,
    ADD PRIMARY KEY (story_id, category_id);
