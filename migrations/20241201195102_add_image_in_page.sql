-- +goose Up

ALTER TABLE story_pages
    ADD COLUMN image_id INT AFTER `page_number`,
    ADD CONSTRAINT fk_page_image
        FOREIGN KEY (image_id)
        REFERENCES images(id)
        ON DELETE CASCADE;

-- +goose Down

ALTER TABLE story_pages
    DROP FOREIGN KEY fk_page_image,
    DROP COLUMN image_id;