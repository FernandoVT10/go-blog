package file

import (
    "mime/multipart"
    "bytes"
    "os"
    "io"
    "path"

    "github.com/h2non/bimg"
)

func SaveFileAsWebp(file multipart.File, filePath string, fileName string) error {
    buffer := bytes.NewBuffer(nil)

    err := os.MkdirAll(filePath, os.ModePerm)
    if err != nil {
        return err
    }

    _, err = io.Copy(buffer, file)
    if err != nil {
        return err
    }

    image := bimg.NewImage(buffer.Bytes())
    imageWebp, err := image.Convert(bimg.WEBP)
    if err != nil {
        return err
    }

    return bimg.Write(path.Join(filePath, fileName), imageWebp)
}
