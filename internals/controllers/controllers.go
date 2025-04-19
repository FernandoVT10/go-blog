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

func saveCover(cover multipart.File) (error, string) {
    coverName := fmt.Sprintf("%d.webp", time.Now().UnixNano())
    err := fileUtils.SaveFileAsWebp(cover, POSTS_UPLOADS_DIR, coverName)

    if err != nil {
        return err, ""
    }

    return nil, coverName
}

func deleteCoverByName(coverName string) error {
    coverPath := path.Join(POSTS_UPLOADS_DIR, coverName)
    return os.Remove(coverPath)
}

func CreateBlogPost(title string, content string, cover multipart.File) (error, string) {
    err, coverName := saveCover(cover)
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
        // TOOD: make this validation part of the validation util
        return errors.New(fmt.Sprintf(`Blog Post with id "%s" not found`, id))
    }

    if err = deleteCoverByName(blogPost.Cover); err != nil {
        return err
    }

    return db.BlogPostModel.DeleteOneById(idInterface)
}

type UpdateBlogPostData struct {
    Title string
    Content string
    Cover multipart.File
}

func UpdateBlogPost(id string, data UpdateBlogPostData) error {
    idObj, err := bson.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    var blogPost db.BlogPost
    db.BlogPostModel.FindById(&blogPost, idObj)

    if blogPost.Id.IsZero() {
        // TOOD: make this validation part of the validation util
        return errors.New(fmt.Sprintf(`Blog Post with id "%s" not found`, id))
    }

    updateData := bson.M{}

    if data.Cover != nil {
        if err = deleteCoverByName(blogPost.Cover); err != nil {
            return err
        }

        err, coverName := saveCover(data.Cover)
        if err != nil {
            return err
        }

        updateData["cover"] = coverName
    }

    if len(data.Title) > 0 {
        updateData["title"] = data.Title
    }
    if len(data.Content) > 0 {
        updateData["content"] = data.Content
    }

    db.BlogPostModel.UpdateById(idObj, updateData)

    return nil
}
