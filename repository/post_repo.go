package repository

import (
	"blog/database"
	"blog/models"
)

type PostRepository interface {
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id string) (*models.Post, error)
	CreatePost(post *models.Post) error
	SavePost(post *models.Post) error
}

type PostRepo struct {
}

func (r *PostRepo) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := database.Db.Preload("Author").Find(&posts).Error
	return posts, err
	
}

func (r *PostRepo) GetPostByID(id string) (*models.Post, error) {
	var post models.Post
	err := database.Db.Preload("Author").Where("ID = ?", id).First(&post).Error
	return &post, err
}

func (r *PostRepo) CreatePost(post *models.Post) error {
	err := database.Db.Create(post).Error
	return err
}

func (r *PostRepo) SavePost(post *models.Post) error {
	err := database.Db.Save(post).Error
	return err
}