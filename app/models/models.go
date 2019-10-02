package models

// Blog model
type Blog struct {
	// ID      bson.ObjectId `json:"_id" bson:"_id"`
	Title   string `json:"title" bson:"title"`
	Content string `json:"content" bson:"content"`
	Author  string `json:"author" bson:"author"`
}

// Response model
type Response struct {
	Message string `json:"message" bson:"message"`
}
