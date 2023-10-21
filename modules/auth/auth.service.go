package auth

import "explore-gofiber/modules/admin"

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

	println("admin", admin)

	if err != nil {
		return result, err
	}

	return result, nil
}
