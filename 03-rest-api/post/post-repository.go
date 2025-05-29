package post

import (
	"fmt"
)

type Post struct {
	Id int `json:"id"`
	PostData
}

type Database struct {
	sequenceNumber int
	posts          []Post
}

func (db *Database) Index() []Post {
	return db.posts
}

func (db *Database) Show(id int) (Post, error) {
	for _, v := range db.posts {
		if v.Id == id {
			return v, nil
		}
	}

	return Post{}, fmt.Errorf("No post with id %d found", id)
}

func (db *Database) Create(data PostData) Post {
	newPost := Post{
		Id:       db.sequenceNumber,
		PostData: data,
	}

	db.sequenceNumber += 1

	db.posts = append(db.posts, newPost)

	return newPost
}

func (db *Database) Delete(id int) error {
	filtered := []Post{}

	for _, v := range db.posts {
		if v.Id != id {
			filtered = append(filtered, v)
		}
	}

	db.posts = filtered

	return nil
}

func NewPostRepository() *Database {
	return &Database{
		sequenceNumber: 1,
		posts:          []Post{},
	}
}
