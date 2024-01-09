package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-project-iman/internal/post-collector-service/app"
	"test-project-iman/internal/post-collector-service/domain"
)

type PostController struct {
	postUseCase app.PostService
}

func NewPostController(postUs app.PostService) *PostController {
	return &PostController{postUseCase: postUs}
}

func (p *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	var post domain.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		http.Error(w, "Invalid Json", http.StatusBadRequest)
	}

	err = p.postUseCase.CollectPosts()
	if err != nil {
		http.Error(w, "Failed to get posts ", http.StatusInternalServerError)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Posts sucessfully fetched and saved to database")
}
