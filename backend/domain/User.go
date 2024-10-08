package domain

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	//IDより前に新たな要素を追加するとUserInfoUpdateが機能しない
	ID           int
	StudentID    string
	Nickname     string
	Email        string
	Password     string
	DepartmentID int
	CourseID     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// マップ型 db を定義
type dbmap map[string]string

// マップの初期化
var DatabaseFields = dbmap{
	"ID":           "id",
	"StudentID":    "student_id",
	"Nickname":     "nickname",
	"Email":        "email",
	"Password":     "password",
	"DepartmentID": "department_id",
	"CourseID":     "course_id",
	"CreatedAt":    "created_at",
	"UpdatedAt":    "updated_at",
}

func (u *User) GeneratePass() error {

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)

	// byteのスライスを返す
	return nil
}
