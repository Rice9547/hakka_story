-- +goose Up

CREATE TABLE images
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    image_url   VARCHAR(255) NOT NULL
);

ALTER TABLE stories
    ADD COLUMN image_id INT,
    ADD CONSTRAINT fk_stories_image
        FOREIGN KEY (image_id)
        REFERENCES images(id);

-- +goose Down

ALTER TABLE stories
    DROP FOREIGN KEY fk_stories_image,
    DROP COLUMN image_id;

DROP TABLE IF EXISTS images;
