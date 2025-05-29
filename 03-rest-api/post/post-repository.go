package post

import (
	"fmt"
)

type Post struct {
	Id int `json:"id"`
	PostData
}

type PostRepository struct {
	sequenceNumber int
	posts          []Post
}

func (db *PostRepository) Index() []Post {
	return db.posts
}

func (db *PostRepository) Show(id int) (Post, error) {
	for _, v := range db.posts {
		if v.Id == id {
			return v, nil
		}
	}

	return Post{}, fmt.Errorf("No post with id %d found", id)
}

func (db *PostRepository) Create(data PostData) Post {
	newPost := Post{
		Id:       db.sequenceNumber,
		PostData: data,
	}

	db.sequenceNumber += 1

	db.posts = append(db.posts, newPost)

	return newPost
}

func (db *PostRepository) Delete(id int) error {
	filtered := []Post{}

	for _, v := range db.posts {
		if v.Id != id {
			filtered = append(filtered, v)
		}
	}

	db.posts = filtered

	return nil
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		sequenceNumber: 1,
		posts:          []Post{},
	}
}
