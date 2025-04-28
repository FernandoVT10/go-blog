package http

import (
    "strings"
    "net/http"
    "fmt"
    "errors"
    "slices"
)

type RequestType int

const (
    Multipart RequestType = 0
)

type Validator interface {
    Validate(r *http.Request) error
    GetKey() string
}

type StringValidator struct {
    Required bool
    MaxLength int
    Key string
}

func (v StringValidator) Validate(r *http.Request) error {
    value := r.FormValue(v.Key)

    if v.Required && value == "" {
        return errors.New(fmt.Sprintf("%s is required", v.Key))
    }

    if len(value) > v.MaxLength {
        return errors.New(fmt.Sprintf("%s must contain %d or less characters", v.Key, v.MaxLength))
    }

    return nil
}
func (v StringValidator) GetKey() string { return v.Key }

type ImageValidator struct {
    Required bool
    Key string
}

var supportedImagesTypes = []string{
    "image/jpg",
    "image/jpeg",
    "image/png",
}

func (v ImageValidator) Validate(r *http.Request) error {
    _, info, err := r.FormFile(v.Key)

    if err != nil {
        if !v.Required {
            return nil
        }

        return errors.New(fmt.Sprintf("%s is required", v.Key))
    }

    fileType := info.Header["Content-Type"][0]

    if !slices.Contains(supportedImagesTypes, fileType) {
        return errors.New("Only jpg, jpeg, and png files are supported")
    }

    // TODO: validate if the image is valid

    return nil
}
func (v ImageValidator) GetKey() string { return v.Key }

func ValidateReq(requestType RequestType, validators []Validator, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vErrors := make(map[string]string)

        switch requestType {
        case Multipart:
            if !strings.Contains(r.Header.Get("Content-Type"), "multipart/form-data") {
                http.Error(w, `Invalid "Content-Type"`, http.StatusBadRequest)
                return
            }
        }

        for _, validator := range validators {
            if err := validator.Validate(r); err != nil {
                vErrors[validator.GetKey()] = err.Error()
            }
        }

        if len(vErrors) > 0 {
            SendJson(w, http.StatusBadRequest, vErrors)
        } else {
            next.ServeHTTP(w, r)
        }
    }
}
