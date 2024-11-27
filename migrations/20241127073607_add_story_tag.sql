-- +goose Up

CREATE TABLE categories (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    name        VARCHAR(255) NOT NULL
);

CREATE TABLE story_to_category (
    story_id    INT,
    category_id INT,
    PRIMARY KEY (story_id, category_id),
    FOREIGN KEY (story_id) REFERENCES stories (id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE IF EXISTS story_to_category;
DROP TABLE IF EXISTS categories;
