package service

import (
	"golang-unit-test/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockRepository struct {
	mock.Mock
}

func (mockRepository *mockRepository) Save(post *entity.Post) error {
	args := mockRepository.Called()
	return args.Error(1)
}

func (mockRepository *mockRepository) FindAll() ([]entity.Post, error) {
	args := mockRepository.Called()
	result := args.Get(0)	
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepository := new(mockRepository)

	post := entity.Post{
		ID: 1,
		Title: "Title",
		Content: "Content",
	}

	// Setup expectations
	mockRepository.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepository)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepository.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "Title", result[0].Title)
	assert.Equal(t, "Content", result[0].Content)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "This post is empty")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{
		ID: 1,
		Title: "",
		Content: "Content",
	}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "This post title is empty")
}
