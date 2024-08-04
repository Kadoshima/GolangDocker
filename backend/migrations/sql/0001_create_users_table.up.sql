-- create_users_table.up.sql
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    student_id VARCHAR(7) NOT NULL,
    nickname VARCHAR(30) NOT NULL,
    email VARCHAR(100) NOT NULL, -- 大学のメアドに限定
    password VARCHAR(512) NOT NULL, -- ハッシュ化
    department_id INT,
    course_id INT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
