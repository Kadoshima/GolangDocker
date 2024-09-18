package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"backend/infrastructure/auth"
)

type CourseUseCaseImpl struct {
	CourseRepository repository.CourseRepository
	JWTService       *auth.JWTService
}

func NewCourseUseCase(repository repository.CourseRepository, jwtService *auth.JWTService) *CourseUseCaseImpl {
	courseUseCaseImpl := &CourseUseCaseImpl{
		CourseRepository: repository,
		JWTService:       jwtService,
	}
	return courseUseCaseImpl
}

func (ci CourseUseCaseImpl) GetAllCourseInfo() ([]*domain.Course, error) {
	// 全てのcourse情報を返す
	courses, err := ci.CourseRepository.SelectAllCourse()
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (ci CourseUseCaseImpl) GetCourseInfo(courseID int) (*domain.Course, error) {
	// 特定のcourse情報を返す
	course, err := ci.CourseRepository.SelectCourse(courseID)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (ci CourseUseCaseImpl) StoreCourseInfo(Course *domain.Course) error {
	// Courseの追加は一旦いらない
	return nil
}
