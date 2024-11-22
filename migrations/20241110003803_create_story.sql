-- +goose Up

CREATE TABLE stories
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    created_at  DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  DATETIME
);

CREATE TABLE story_pages
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    story_id      INT NOT NULL,
    page_number   INT NOT NULL,
    content_cn    TEXT,
    content_hakka TEXT,
    created_at    DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at    DATETIME,
    FOREIGN KEY (story_id) REFERENCES stories (id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE IF EXISTS story_pages;
DROP TABLE IF EXISTS stories;
