package database

import "rest-api/data"

type Post struct {
	Id int `json:"id"`
	data.PostData
}

type Database struct {
	sequenceNumber int
	posts          []Post
}

func (db *Database) List() []Post {
	return db.posts
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
