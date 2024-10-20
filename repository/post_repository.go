package repository

import (
	"golang-unit-test/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Save(post *entity.Post) error
	FindAll() ([]entity.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (postRepository *postRepository) Save(post *entity.Post) error {
	return postRepository.db.Create(post).Error
}

func (postRepository *postRepository) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	err := postRepository.db.Find(&posts).Error
	return posts, err
}
