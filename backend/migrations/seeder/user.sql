-- User テーブルへの INSERT クエリ
INSERT INTO users (id, student_id, nickname, email, password, department_id, course_id, created_at, updated_at)
VALUES
    (1, 'S123456', 'JohnDoe', 'john.doe@example.com', '$2a$10$Z7j9LwHV7JgFWP8vLWvWyeSRv3gF2K1u/TP6vYdrKKUzrZ68FiK5i', 101, 201, '2024-09-16 08:00:00', '2024-09-16 08:00:00'),
    (2, 'S789012', 'JaneSmith', 'jane.smith@example.com', '$2a$10$6FfRtFhWUl/tlFs2SLJ9H.dD1G9HqE6cG9Y1OuFk5Pj2Afri/cOqG', 102, 202, '2024-09-16 08:05:00', '2024-09-16 08:05:00'),
    (3, 'S345678', 'BobJohnson', 'bob.johnson@example.com', '$2a$10$6Fhg2I8N/OBz5zSLFryGhOV.zpKlF/8KMkGf2i.kFT2kC8uR5/x/y', 103, 203, '2024-09-16 08:10:00', '2024-09-16 08:10:00');
