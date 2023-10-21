package auth

import (
	"explore-gofiber/modules/admin"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type IService interface {
	Login(dto LoginDto) (LoginResult, error)
}

type service struct {
	adminService admin.IService
}

func NewService(adminService admin.IService) *service {
	return &service{
		adminService,
	}
}

func (s *service) Login(dto LoginDto) (LoginResult, error) {
	var result LoginResult

	admin, err := s.adminService.FindByEmail(dto.Email)

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

	return result, nil
}
