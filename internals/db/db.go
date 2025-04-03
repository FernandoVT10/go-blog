package db

import (
    "time"
    "context"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
)

const MONGO_URI = "mongodb://localhost:27017/"
const DB_TIMEOUT = 500 * time.Millisecond

var db *mongo.Database

func Connect() {
    client, err := mongo.Connect(options.Client().ApplyURI(MONGO_URI))
    if err != nil {
        panic(err)
    }

    db = client.Database("go-blog")
}

type Model struct {
    collectionName string
}

func (model Model) FindOne(result interface{}) error {
    ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
    defer cancel()

    coll := db.Collection(model.collectionName)

    return coll.FindOne(ctx, bson.D{{"title", "Test"}}).Decode(result)
}

func (model Model) Find(result interface{}) error {
    ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
    defer cancel()

    coll := db.Collection(model.collectionName)
    cur, err := coll.Find(ctx, bson.M{})

    if err != nil {
        return err
    }

    return cur.All(context.Background(), result)
}

var BlogPostModel = Model{
    collectionName: "BlogPosts",
}

type BlogPost struct {
    Id bson.ObjectID `bson:"_id,omitempty"`
    Title string `bson:"title,omitempty"`
    Cover string `bson:"cover,omitempty"`
    Content string `bson:"content,omitempty"`
}
