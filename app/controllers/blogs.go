package controllers

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/anirudh06/blog-api-with-golang/app/models"
)

// PostCreateBlog is for creating blog
func PostCreateBlog(db *mongo.Database, res http.ResponseWriter, req *http.Request) {
	blog := models.Blog{}
	blog.Title = req.FormValue("title")
	blog.Content = req.FormValue("content")
	blog.Author = req.FormValue("author")
	_, err := db.Collection("blogs").InsertOne(context.TODO(), blog)
	if err != nil {
		log.Fatal(err)
	}
	respondJSON(res, http.StatusOK, blog)
}
