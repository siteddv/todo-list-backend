package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository"
)

const (
	// salt contains key for generating password hash
	salt = "jdkshf54jh5jh34kg4weknjf"

	// signingKey contains key for decrypt and encrypt token
	signingKey = "sjkfkjha234uilweorit34"

	// tokenTTL is time of life-cycle of generated token
	tokenTTL = 12 * time.Hour
)

// tokenClaims is a custom Claims containing jwt.StandardClaims and id of signed-in user
type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

// AuthService contains repository for working with auth in db
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService returns pointer on a new instance of AuthService
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser creates model.User in DB using specified user model. It returns new user id and error
func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	userId, err := s.repo.CreateUser(user)

	return userId, err
}

// GenerateToken generate token for signing in user. Returns a complete token and error
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	password = generatePasswordHash(password)
	user, err := s.repo.GetUser(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	signedToken, err := token.SignedString([]byte(signingKey))

	return signedToken, err
}

// ParseToken decrypts token and returns id of signed in user and error
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	parsingFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	}

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, parsingFunc)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

// generatePasswordHash generates password hash using SHA-1 algorithm
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	hashedPassword := fmt.Sprintf("%x", hash.Sum([]byte(salt)))

	return hashedPassword
}
