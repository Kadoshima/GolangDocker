-- 0004_create_forums_table.up.sql

CREATE TABLE IF NOT EXISTS forums (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_by INT NOT NULL,
    status INT NOT NULL,
    visibility INT NOT NULL,
    category VARCHAR(100),
    num_posts INT DEFAULT 0,
    attachments VARCHAR(512),
    moderators JSON,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users(id)
);
