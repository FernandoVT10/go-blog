package controllers

import (
    "fmt"
    "errors"
    "mime/multipart"
    "time"
    "path"
    "os"

    "github.com/FernandoVT10/go-blog/internals/db"
    "go.mongodb.org/mongo-driver/v2/bson"

    fileUtils "github.com/FernandoVT10/go-blog/internals/utils/file"
)

const POSTS_UPLOADS_DIR = "./uploads/posts"

func CreateBlogPost(title string, content string, cover multipart.File) (error, string) {
    coverName := fmt.Sprintf("%d.webp", time.Now().UnixNano())
    err := fileUtils.SaveFileAsWebp(cover, POSTS_UPLOADS_DIR, coverName)

    if err != nil {
        return err, ""
    }

    err, id := db.BlogPostModel.CreateOne(bson.M{
        "title": title,
        "content": content,
        "cover": coverName,
    })

    return nil, id.(bson.ObjectID).Hex()
}

func DeleteBlogPost(id string) error {
    idInterface, err := bson.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    var blogPost db.BlogPost
    db.BlogPostModel.FindById(&blogPost, idInterface)

    if blogPost.Id.IsZero() {
        return errors.New(fmt.Sprintf(`Blog Post with id "%s" not found`, id))
    }

    coverPath := path.Join(POSTS_UPLOADS_DIR, blogPost.Cover)
    err = os.Remove(coverPath)
    if err != nil {
        return err
    }

    return db.BlogPostModel.DeleteOneById(idInterface)
}
