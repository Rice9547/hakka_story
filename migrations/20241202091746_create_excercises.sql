-- +goose Up

CREATE TABLE `exercises` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `story_id` INT NOT NULL,
    `type` TINYINT NOT NULL COMMENT '1: Text and Audio with Open Answer, 2: Multiple Choice',
    `prompt_text` TEXT NOT NULL COMMENT 'Text prompt for the exercise',
    `audio_url` VARCHAR(255) DEFAULT NULL COMMENT 'Optional audio prompt',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`story_id`) REFERENCES `stories`(`id`) ON DELETE CASCADE
);

CREATE TABLE `exercise_open_answers` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `exercise_id` INT NOT NULL,
    `answer_text` VARCHAR(255) NOT NULL COMMENT 'Possible correct answer',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`exercise_id`) REFERENCES `exercises`(`id`) ON DELETE CASCADE
);

CREATE TABLE `exercise_choices` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `exercise_id` INT NOT NULL,
    `choice_text` VARCHAR(255) NOT NULL COMMENT 'Choice text',
    `is_correct` BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'Whether this choice is part of the correct answer',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (`exercise_id`) REFERENCES `exercises`(`id`) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE `exercise_choices`;
DROP TABLE `exercise_open_answers`;
DROP TABLE `exercises`;
