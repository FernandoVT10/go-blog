package controllers

import (
    "net/http"
    "errors"
    "github.com/golang-jwt/jwt/v5"
)

// TODO: get this key from a file or envfile
const JWT_SECRET = "hello"

func Login(password string) (string, error) {
    if password != "secret" {
        return "", errors.New("Incorrect Password")
    }

    t := jwt.New(jwt.SigningMethodHS512)
    signedStr, err := t.SignedString([]byte(JWT_SECRET))
    if err != nil {
        println(err.Error())
        return "", errors.New("Server Error")
    }

    return signedStr, nil
}

// returns true if the token is valid
func validateToken(token string) bool {
    parser := jwt.NewParser(
        jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Name}),
    )

    _, err := parser.Parse(token, func(_ *jwt.Token) (any, error) {
        return []byte(JWT_SECRET), nil
    })

    return err == nil
}

func IsAuthenticated(r *http.Request) bool {
    cookie, err := r.Cookie("token")
    if err != nil {
        return false
    }

    token := cookie.Value
    return validateToken(token)
}
