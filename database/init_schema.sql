CREATE TABLE IF NOT EXISTS videos (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    video_id VARCHAR(255) UNIQUE KEY,
    title TEXT NOT NULL,
    description TEXT,
    thumbnail_url VARCHAR(255) NOT NULL,
    publish_time DATETIME NOT NULL,
    FULLTEXT title_description_idx (title, description)
);