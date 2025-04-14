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

const DescendingSort = -1;
const AscendingSort = 1;

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
    useTimestamps bool
}

func (model Model) FindOne(result any, filter bson.D) error {
    ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
    defer cancel()

    coll := db.Collection(model.collectionName)

    return coll.FindOne(ctx, filter).Decode(result)
}

func (model Model) FindById(result any, id bson.ObjectID) error {
    ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
    defer cancel()

    coll := db.Collection(model.collectionName)

    return coll.FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(result)
}

func (model Model) Find(result any, filter bson.D, opts *options.FindOptionsBuilder) error {
    ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
    defer cancel()

    coll := db.Collection(model.collectionName)
    cur, err := coll.Find(ctx, filter, opts)

    if err != nil {
        return err
    }

    return cur.All(context.Background(), result)
}

func (model Model) CreateOne(data bson.M) (error, any) {
    ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
    defer cancel()

    if model.useTimestamps {
        data["createdAt"] = time.Now()
    }
    coll := db.Collection(model.collectionName)

    res, err := coll.InsertOne(ctx, data)
    if err != nil {
        return err, nil
    }

    return nil, res.InsertedID
}

func (model Model) DeleteOneById(id bson.ObjectID) error {
    ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
    defer cancel()

    coll := db.Collection(model.collectionName)
    _, err := coll.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
    if err != nil {
        return err
    }

    return nil
}

var BlogPostModel = Model{
    collectionName: "BlogPosts",
    useTimestamps: true,
}

type BlogPost struct {
    Id bson.ObjectID `bson:"_id,omitempty"`
    Title string `bson:"title,omitempty"`
    Cover string `bson:"cover,omitempty"`
    Content string `bson:"content,omitempty"`
    CreatedAt time.Time `bson:"createdAt"`
}
