package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
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

func GeneratePass(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	// byteのスライスを返す
	return hash, nil
}
