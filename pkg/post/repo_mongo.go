package post

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrNotFound = errors.New("Not found")
)

type PostsMongo struct {
	DB *mgo.Collection
}

func NewPostsMongo() *PostsMongo {
	mongoSession, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		fmt.Println("err")
		panic(err)
	}
	collectionPosts := mongoSession.DB("reddit").C("posts")
	if n, _ := collectionPosts.Count(); n == 0 {
		collectionPosts.Insert(&Post{Id: bson.NewObjectId(),
			Author:   Author{Username: "fdsd"},
			Category: "music",
			Text:     "sdavdfdfsvsdfv",
			Title:    "sfdfdsfds",
			Type:     "text"})
	}
	return &PostsMongo{DB: collectionPosts}
}

func (bd *PostsMongo) GetAll() ([]*Post, error) {
	posts := []*Post{}

	err := bd.DB.Find(bson.M{}).All(&posts)
	if err != nil {
		fmt.Errorf("bd can't find: %w", err)
		return nil, err
	}

	return posts, nil
}

func (bd *PostsMongo) GetByCategory(category string) ([]*Post, error) {
	posts := []*Post{}

	err := bd.DB.Find(bson.M{"category": category}).All(&posts)
	if err != nil {
		fmt.Errorf("bd can't find: %w", err)
		return nil, err
	}

	return posts, nil
}

func (bd *PostsMongo) GetByUser(username string) ([]*Post, error) {
	posts := []*Post{}

	err := bd.DB.Find(bson.M{"author.username": username}).All(&posts)
	if err != nil {
		fmt.Errorf("bd can't find: %w", err)
		return nil, err
	}

	return posts, nil
}

func (bd *PostsMongo) GetById(id string) (*Post, error) { // добавить views
	post := Post{}

	if !bson.IsObjectIdHex(id) {
		return nil, ErrNotFound
	}

	bsonId := bson.ObjectIdHex(id)
	err := bd.DB.Find(bson.M{"id": bsonId}).One(&post)
	if err != nil {
		fmt.Errorf("bd can't find: %w", err)
		return nil, err
	}

	return &post, nil
}

func (bd *PostsMongo) AddPost(post *Post) (string, error) {
	err := bd.DB.Insert(post)
	if err != nil {
		fmt.Errorf("bd can't add: %w", err)
		return "", err
	}
	return string(post.Id), nil
}

func (bd *PostsMongo) UpdatePost(post *Post) (*Post, error) {

	err := bd.DB.Update(bson.M{"id": post.Id}, post)
	if err != nil {
		fmt.Errorf("bd can't update: %w", err)
		return nil, err
	}

	return post, nil
}

func (bd *PostsMongo) Delete(id string) (bool, error) {
	if !bson.IsObjectIdHex(id) {
		return false, ErrNotFound
	}

	bsonId := bson.ObjectIdHex(id)
	err := bd.DB.Remove(bson.M{"id": bsonId})
	if err != nil {
		fmt.Errorf("bd can't delete: %w", err)
		return false, err
	}
	return true, nil
}
