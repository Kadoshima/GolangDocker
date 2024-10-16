package usecase

import (
	"backend/adapter/repository"
	"backend/infrastructure/auth"
	"backend/service"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCaseImpl struct {
	AuthRepository repository.AuthRepository
	MinioService   service.MinioService
	JWTService     *auth.JWTService
}

func NewAuthUseCase(authRepository repository.AuthRepository, minioService service.MinioService, jwtService *auth.JWTService) AuthUseCase {
	return &AuthUseCaseImpl{
		AuthRepository: authRepository,
		MinioService:   minioService,
		JWTService:     jwtService,
	}
}

func (au *AuthUseCaseImpl) Login(userID int, password string) (string, error) {

	// パスワードを取得
	hashedPassword, err := au.AuthRepository.GetPasswordByUserID(userID)
	if err != nil {
		return "", errors.New("could not retrieve user password")
	}

	// パスワードの検証
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	// tokenの発行
	token, err := au.JWTService.GenerateJWT(userID)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}

func (au *AuthUseCaseImpl) GetUserIDByStudentID(studentID string) (int, error) {

	// studentIDからuserIDの取得
	userID, err := au.AuthRepository.GetUserIDByStudentID(studentID)
	if err != nil {
		return -1, errors.New("could not retrieve user password")
	}

	return userID, nil
}
