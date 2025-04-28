package controllers

import (
    "net/http"
    "errors"
    "github.com/golang-jwt/jwt/v5"

    "github.com/FernandoVT10/go-blog/app/config"
)

func Login(password string) (string, error) {
    if password != config.GetEnv().AdminPass {
        return "", errors.New("Incorrect Password")
    }

    t := jwt.New(jwt.SigningMethodHS512)
    signedStr, err := t.SignedString(config.GetEnv().JwtSecret)
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
        return config.GetEnv().JwtSecret, nil
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
