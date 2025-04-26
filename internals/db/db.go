package db

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/mongo/readpref"

    "github.com/FernandoVT10/go-blog/internals/config"
)

const DB_TIMEOUT = 500 * time.Millisecond

const DescendingSort = -1;
const AscendingSort = 1;

var db *mongo.Database

func Connect() {
    fmt.Println("[INFO] Connecting to mongo database...")
    client, err := mongo.Connect(options.Client().ApplyURI(config.GetEnv().MongoUri))
    if err != nil {
        panic(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    err = client.Ping(ctx, readpref.Primary())

    if err != nil {
        panic(err)
    }

    fmt.Println("[INFO] Database connected successfully")

    db = client.Database(config.GetEnv().MongoDbName)
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

func (model Model) UpdateById(id bson.ObjectID, data bson.M) error {
    ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
    defer cancel()

    coll := db.Collection(model.collectionName)

    _, err := coll.UpdateByID(ctx, id, bson.M{ "$set": data })
    return err
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
