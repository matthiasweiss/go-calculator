package database

import (
	"fmt"
	"rest-api/data"
)

type Post struct {
	Id int `json:"id"`
	data.PostData
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

func (db *Database) Create(data data.PostData) Post {
	newPost := Post{
		Id:       db.sequenceNumber,
		PostData: data,
	}

	db.sequenceNumber += 1

	db.posts = append(db.posts, newPost)

	return newPost
}

func NewDatabase() *Database {
	return &Database{
		sequenceNumber: 1,
		posts:          []Post{},
	}
}
