package services

import (
	"blog/database"
	"blog/models"
	"blog/repository"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type PostService struct {
	Repo repository.PostRepository
}

func (s *PostService) ListAllPosts() ([]models.Post, error) {
	posts, err := s.Repo.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GetPost(id string) (*models.Post, error) {
	post, err := s.Repo.GetPostByID(id)
	if err != nil {
		return &models.Post{}, err
	}
	return post, nil
}

func (s *PostService) CreatePost(req *models.Post, claims jwt.MapClaims) error {
	userId, err := uuid.Parse(claims["userID"].(string))
	if err != nil {
		return err
	}
	req.AuthorID = userId
	err = s.Repo.CreatePost(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) UpdatePost(req *models.Post, postID string, claims jwt.MapClaims) error {
	post, err := s.Repo.GetPostByID(postID)
	if err != nil {
		return err
	}
	userId, err := uuid.Parse(claims["userID"].(string))
	if err != nil {
		return err
	}
	if post.AuthorID != userId {
		return errors.New("cannot edit others posts")
	}

	post.Title = req.Title
	post.Content = req.Content
	req = post
	err = s.Repo.SavePost(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) DeletePost(claims jwt.MapClaims, id string) error {
	post, err := s.Repo.GetPostByID(id)
	if err != nil {
		return err
	}
	userId, err := uuid.Parse(claims["userID"].(string))
	if err != nil {
		return err
	}
	if post.AuthorID != userId {
		return errors.New("cannot delete others posts")
	}

	err = database.Db.Delete(post).Error
	if err != nil {
		return err
	}
	return nil
}