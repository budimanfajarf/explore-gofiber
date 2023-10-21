package auth

import (
	"explore-gofiber/modules/admin"
	"explore-gofiber/utils/jwt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Login(dto LoginDto) (AuthResponse, error)
}

type service struct {
	adminRepository admin.IRepository
}

func NewService(adminRepository admin.IRepository) *service {
	return &service{
		adminRepository,
	}
}

func (s *service) Login(dto LoginDto) (AuthResponse, error) {
	var result AuthResponse

	admin := &UserResponse{}

	err := s.adminRepository.FindOneByEmail(admin, dto.Email).Error

	invalidErrMsg := "invalid email or password"

	if err != nil {
		if err.Error() == "record not found" {
			return result, fiber.NewError(fiber.StatusBadRequest, invalidErrMsg)
		}

		return result, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(dto.Password))

	if err != nil {
		println(err.Error())
		return result, fiber.NewError(fiber.StatusBadRequest, invalidErrMsg)
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: admin.ID,
	})

	result = AuthResponse{
		User: admin,
		Auth: &AccessResponse{
			Token: token,
		},
	}

	return result, nil
}
