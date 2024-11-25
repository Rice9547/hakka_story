-- +goose Up

CREATE TABLE audio_files (
    id              INT AUTO_INCREMENT PRIMARY KEY,
    story_page_id   INT NOT NULL,
    dialect         VARCHAR(255) NOT NULL,
    audio_url       VARCHAR(255) NOT NULL,
    FOREIGN KEY (story_page_id) REFERENCES story_pages (id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE audio_files;
