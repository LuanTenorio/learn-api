package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/LuanTenorio/learn-api/internal/auth"
	"github.com/LuanTenorio/learn-api/internal/auth/dto"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (uc *authUseCaseImpl) Login(ctx context.Context, loginDto *dto.LoginDTO) (string, error) {
	user, err := uc.userRepository.FindUserAndPwdByEmail(ctx, loginDto.Email)

	if err != nil {
		return "", err
	}

	if err := checkPwd(user.Password, loginDto.Password); err != nil {
		return "", errors.New("wrong password")
	}

	token, err := createToken(user.Name)

	if err != nil {
		return "", err
	}

	return token, nil
}

func checkPwd(hash string, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24 * 20).Unix(),
		})

	return token.SignedString(auth.SecretKey)
}
