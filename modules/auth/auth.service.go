package auth

import (
	"explore-gofiber/modules/admin"
	"explore-gofiber/utils/jwt"
	"explore-gofiber/utils/password"

	"github.com/gofiber/fiber/v2"
)

type IService interface {
	Login(dto LoginDto) (AuthData, error)
}

type service struct {
	adminRepository admin.IRepository
}

func NewService(adminRepository admin.IRepository) *service {
	return &service{
		adminRepository,
	}
}

func (s *service) Login(dto LoginDto) (AuthData, error) {
	var data AuthData

	admin := &AdminData{}

	invalidErrMsg := "invalid email or password"

	err := s.adminRepository.FindOneByEmail(admin, dto.Email).Error
	if err != nil {
		if err.Error() == "record not found" {
			return data, fiber.NewError(fiber.StatusBadRequest, invalidErrMsg)
		}

		return data, err
	}

	err = password.Verify(admin.Password, dto.Password)
	if err != nil {
		return data, fiber.NewError(fiber.StatusBadRequest, invalidErrMsg)
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: admin.ID,
	})

	data = AuthData{
		ID:    admin.ID,
		Role:  admin.Role,
		Name:  admin.Name,
		Email: admin.Email,
		Token: token,
	}

	return data, nil
}
