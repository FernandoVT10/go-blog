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
    "go.mongodb.org/mongo-driver/v2/mongo/options"

    fileUtils "github.com/FernandoVT10/go-blog/internals/utils/file"
)

// TODO: make this configurable
const POSTS_UPLOADS_URL = "http://localhost:3000/uploads/posts"
const POSTS_UPLOADS_DIR = "./uploads/posts"

func ConvertCoverToUrl(cover string) string {
    return fmt.Sprintf("%s/%s", POSTS_UPLOADS_URL, cover)
}

// converts all covers names (19238...2.webp) into an url that can be send to the frontend
func ConvertCoversToUrl(blogPosts []db.BlogPost) {
    for i := range len(blogPosts) {
        blogPosts[i].Cover = ConvertCoverToUrl(blogPosts[i].Cover)
    }
}

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

type GetBlogPostsOpts struct {
    Limit int64
    Sort map[string]int
}

// returns posts with the cover as an url
func GetBlogPosts(opts GetBlogPostsOpts) []db.BlogPost {
    var blogPosts []db.BlogPost

    queryOpts := options.Find()

    if opts.Limit > 0 {
        queryOpts = queryOpts.SetLimit(opts.Limit)
    }


    if len(opts.Sort) > 0 {
        sortOpts := bson.D{}

        for key, sortOpt := range opts.Sort {
            sortOpts = append(sortOpts, bson.E{Key: key, Value: sortOpt})
        }

        queryOpts = queryOpts.SetSort(sortOpts)
    }

    db.BlogPostModel.Find(&blogPosts, bson.D{}, queryOpts)
    ConvertCoversToUrl(blogPosts)

    return blogPosts
}

func GetBlogPostByHexId(id string) (db.BlogPost, error) {
    idObj, err := bson.ObjectIDFromHex(id)
    if err != nil {
        return db.BlogPost{}, nil
    }

    var blogPost db.BlogPost
    err = db.BlogPostModel.FindById(&blogPost, idObj)
    if err != nil {
        return db.BlogPost{}, nil
    }

    blogPost.Cover = ConvertCoverToUrl(blogPost.Cover)

    return blogPost, nil
}
