package repository

import (
	"backend/domain"
	"database/sql"
	"time"
)

type CourseRepositoryImpl struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepositoryImpl {
	return &CourseRepositoryImpl{db: db}
}

// 全てのコースを取得
func (cr *CourseRepositoryImpl) SelectAllCourse() ([]*domain.Course, error) {
	rows, err := cr.db.Query("SELECT id, department_id, name FROM courses ORDER BY department_id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // rows を必ずクローズ

	var courses []*domain.Course // スライスの要素はポインタ型
	for rows.Next() {
		var course domain.Course
		if err := rows.Scan(&course.ID, &course.DepartmentID, &course.Name); err != nil {
			return nil, err
		}
		courses = append(courses, &course) // courseのポインタを追加
	}

	// ループ終了後のエラーチェック
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (cr *CourseRepositoryImpl) SelectCourse(courseID int) (*domain.Course, error) {
	course := &domain.Course{}
	err := cr.db.QueryRow(
		"SELECT id, department_id, name FROM courses WHERE id = ?",
		courseID,
	).Scan(&course.ID, &course.DepartmentID, &course.Name)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (cr *CourseRepositoryImpl) CreateCourse(course *domain.Course) (*domain.Course, error) {
	now := time.Now()
	course.CreatedAt = now
	course.UpdatedAt = now

	result, err := cr.db.Exec(
		"INSERT INTO courses (department_id, name, created_at, updated_at) VALUES (?, ?, ?, ?)",
		course.DepartmentID, course.Name, course.CreatedAt, course.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	courseID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	course.ID = int(courseID)
	return course, nil
}

func (cr *CourseRepositoryImpl) UpdateCourse(course *domain.Course) (*domain.Course, error) {
	course.UpdatedAt = time.Now()
	_, err := cr.db.Exec(
		"UPDATE courses SET department_id = ?, name = ?, updated_at = ? WHERE id = ?",
		course.DepartmentID, course.Name, course.UpdatedAt, course.ID,
	)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (cr *CourseRepositoryImpl) DeleteCourse(courseID int) error {
	_, err := cr.db.Exec(
		"DELETE FROM courses WHERE id = ?",
		courseID,
	)
	return err
}
