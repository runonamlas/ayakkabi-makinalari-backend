package service

import (
	"github.com/mashingan/smapping"
	"github.com/my-way-teams/my_way_backend/dto"
	"github.com/my-way-teams/my_way_backend/entity"
	"github.com/my-way-teams/my_way_backend/repository"
	"log"
)

type AdminService interface {
	VerifyCredential(email string) interface{}
	VerifyPassword(res interface{},password string) interface{}
	Create(admin dto.AdminRegisterDTO) entity.Admin
	Update(admin dto.AdminUpdateDTO) entity.Admin
	Profile(adminID string) entity.Admin
	Users() []entity.User
	FindByEmail(email string) entity.Admin
	IsDuplicateEmail(email string) bool
	FindByUsername(username string) entity.Admin
	IsDuplicateUsername(username string) bool
}

type adminService struct {
	adminRepository repository.AdminRepository
}

func NewAdminService(adminRep repository.AdminRepository) AdminService {
	return &adminService{
		adminRepository: adminRep,
	}
}

func (service *adminService) VerifyCredential(email string) interface{}  {
	res := service.adminRepository.VerifyCredential(email)
	if v, ok := res.(entity.Admin); ok {
		if v.Email == email || v.Username == email {
			return res
		}
		return false
	}
	return false
}

func (service *adminService) VerifyPassword(res interface{}, password string) interface{}  {
	if v, ok := res.(entity.Admin); ok {
		comparedPassword := comparedPassword(v.Password, []byte(password))
		if comparedPassword {
			return res
		}
	}
	return false
}

func (service *adminService) Create(admin dto.AdminRegisterDTO) entity.Admin {
	adminToCreate := entity.Admin{}
	err := smapping.FillStruct(&adminToCreate, smapping.MapFields(&admin))
	if err != nil {
		log.Fatalf("Failed Map %v", err)
	}
	res := service.adminRepository.InsertAdmin(adminToCreate)
	return res
}

func (service *adminService) Update(admin dto.AdminUpdateDTO) entity.Admin {
	adminToUpdate := entity.Admin{}
	err := smapping.FillStruct(&adminToUpdate, smapping.MapFields(&admin))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updatedAdmin := service.adminRepository.UpdateAdmin(adminToUpdate)
	return updatedAdmin
}

func (service *adminService) Profile(adminID string) entity.Admin {
	return service.adminRepository.ProfileAdmin(adminID)
}

func (service *adminService) Users() []entity.User {
	return service.adminRepository.Users()
}


func (service *adminService) FindByEmail(email string) entity.Admin {
	return service.adminRepository.FindByEmail(email)
}

func (service *adminService) IsDuplicateEmail(email string) bool {
	res := service.adminRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *adminService) FindByUsername(username string) entity.Admin {
	return service.adminRepository.FindByUsername(username)
}

func (service *adminService) IsDuplicateUsername(username string) bool {
	res := service.adminRepository.IsDuplicateUsername(username)
	return !(res.Error == nil)
}