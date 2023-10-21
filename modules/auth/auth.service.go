package auth

import (
	"errors"
	"explore-gofiber/constant"
	"explore-gofiber/modules/admin"
	"explore-gofiber/utils/jwt"
	"explore-gofiber/utils/password"
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

	err := s.adminRepository.FindOneByEmail(admin, dto.Email).Error
	if err != nil {
		if err.Error() == "record not found" {
			return data, errors.New(constant.ErrInvalidCredentials)
		}

		return data, err
	}

	err = password.Verify(admin.Password, dto.Password)
	if err != nil {
		return data, errors.New(constant.ErrInvalidCredentials)
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
