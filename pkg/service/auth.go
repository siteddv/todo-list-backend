package service

import (
	"crypto/sha1"
	"fmt"
	todo "todolistBackend"
	"todolistBackend/pkg/repository"
)

const salt = "jdkshf54jh5jh34kg4weknjf"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", hash.Sum([]byte(salt)))

	return hashedPassword
}
