package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/LuanTenorio/learn-api/internal/auth"
	"github.com/LuanTenorio/learn-api/internal/auth/dto"
	userEntity "github.com/LuanTenorio/learn-api/internal/user/entity"
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

	token, err := createToken(&userEntity.User{Id: user.Id, Name: user.Name, Email: user.Email})

	if err != nil {
		return "", err
	}

	return token, nil
}

func checkPwd(hash string, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
}

func createToken(usr *userEntity.User) (string, error) {
	claims := &auth.JwtCustomClaims{
		User: *usr,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 15)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(auth.SecretKey))
}
