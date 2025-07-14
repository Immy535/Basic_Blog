package services

import (
	"blog/middleware"
	"blog/models"
	"blog/repository"
	"blog/utils"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) Register(req *models.User) error {
	_, err := s.Repo.GetUserByEmail(req.Email)
	if err == nil {
		return errors.New("email is linked to an existing user")
	}

	hashPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	req.Password = hashPass

	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(req *models.User) (string, error) {
	user, err := s.Repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *UserService) LoginInfo(claims jwt.MapClaims) (*models.User,error) {
	userId := claims["userID"].(string)
	user, err := s.Repo.GetUserByID(userId)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}