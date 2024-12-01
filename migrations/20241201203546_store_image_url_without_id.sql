-- +goose Up

ALTER TABLE stories
    ADD COLUMN image_url VARCHAR(255) NOT NULL DEFAULT '' AFTER description;

UPDATE stories
    JOIN images ON stories.image_id = images.id
    SET stories.image_url = images.image_url;

ALTER TABLE stories
    DROP FOREIGN KEY fk_stories_image,
    DROP COLUMN image_id;

ALTER TABLE story_pages
    ADD COLUMN image_url VARCHAR(255) NOT NULL DEFAULT '' AFTER page_number;

UPDATE story_pages
    JOIN images ON story_pages.image_id = images.id
    SET story_pages.image_url = images.image_url;

ALTER TABLE story_pages
    DROP FOREIGN KEY fk_page_image,
    DROP COLUMN image_id;

DROP TABLE images;

-- +goose Down

CREATE TABLE images (
    id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    image_url VARCHAR(255) NOT NULL DEFAULT ''
);

INSERT INTO images (image_url)
    SELECT DISTINCT image_url
    FROM (
        SELECT image_url FROM story_pages
        UNION
        SELECT image_url FROM stories
    ) AS combined_urls;

ALTER TABLE story_pages
    ADD COLUMN image_id INT AFTER page_number;

UPDATE story_pages
    JOIN images ON story_pages.image_url = images.image_url
    SET story_pages.image_id = images.id;

ALTER TABLE story_pages
    DROP COLUMN image_url;

ALTER TABLE stories
    ADD COLUMN image_id INT AFTER description;

UPDATE stories
    JOIN images ON stories.image_url = images.image_url
    SET stories.image_id = images.id;

ALTER TABLE stories
    DROP COLUMN image_url;
