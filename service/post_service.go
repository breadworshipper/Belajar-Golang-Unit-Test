package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{
		repository: postRepository,
	}
}

func (postService *postService) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("This post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("This post title is empty")
		return err
	}

	return nil
}

func (postService *postService) Create(post *entity.Post) (*entity.Post, error) {
	err := postService.repository.Save(post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (postService *postService) FindAll() ([]entity.Post, error) {
	result, err := postService.repository.FindAll()
	if err != nil {
		return result,  errors.New("Error finding all posts")
	}
	
	return result, nil
}
