package handler

import (
	"golang-unit-test/entity"
	"golang-unit-test/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type postHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) postHandler {
	return postHandler{
		postService: postService,
	}
}

func (postHandler *postHandler) PostCreateHandler(context *gin.Context) {
    var post entity.Post

    if err := context.ShouldBindJSON(&post); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := postHandler.postService.Create(&post)
    if err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    context.JSON(http.StatusCreated, gin.H{"data": post})
}
