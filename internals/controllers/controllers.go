package controllers

import (
    "errors"

    "github.com/golang-jwt/jwt/v5"
)

func Login(password string) (string, error) {
    if password != "secret" {
        return "", errors.New("Incorrect Password")
    }

    t := jwt.New(jwt.SigningMethodHS512)
    signedStr, err := t.SignedString([]byte("hello"))
    if err != nil {
        println(err.Error())
        return "", errors.New("Server Error")
    }

    return signedStr, nil
}
