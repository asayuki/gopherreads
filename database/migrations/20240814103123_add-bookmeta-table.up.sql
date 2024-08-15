CREATE TABLE IF NOT EXISTS bookmeta(
    `book_id` INTEGER,
    `title` VARCHAR(255) NOT NULL DEFAULT '',
    `description` TEXT NOT NULL DEFAULT '',
    `author` VARCHAR(255) NOT NULL DEFAULT '',
    `genre` VARCHAR(255) NOT NULL DEFAULT '',
    `cover` VARCHAR(100) NOT NULL DEFAULT '',
    FOREIGN KEY(`book_id`) REFERENCES books(`id`)
)