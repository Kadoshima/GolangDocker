-- 0005_create_support_request_table.up.sql

CREATE TABLE IF NOT EXISTS support_requests (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT 'サポート要請ID',
    forum_id INT NOT NULL COMMENT '掲示板ID',
    post_id INT NOT NULL COMMENT '投稿ID',
    request_content VARCHAR(255) NOT NULL COMMENT '要請内容',
    request_department tinyint(4) COMMENT '答えてほしい学部',
    progress_status tinyint(4) NOT NULL COMMENT '進行レベル',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
    CONSTRAINT fk_forum_id FOREIGN KEY (forum_id) REFERENCES forums(id),
    CONSTRAINT fk_post_id FOREIGN KEY (post_id) REFERENCES posts(id)
);
