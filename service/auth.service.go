package service

import (
	"log"

	"github.com/mashingan/smapping"
	"golang-api/dto"
	"golang-api/models"
	"golang-api/repository"
	"golang.org/x/crypto/bcrypt"
)

// AuthService is a contract about something that this service can do
type AuthService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email, password string) interface{} {

	res := service.userRepository.VerifyCredential(email, password)

	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}

	return res
}

func (service *authService) CreateUser(user dto.RegisterDTO) models.User {
	userToCreate := models.User{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.userRepository.InsertUser(userToCreate)
	return res
}

func (service *authService) FindByEmail(email string) models.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPwd []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), plainPwd)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
