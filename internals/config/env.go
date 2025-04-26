package config

import (
    "fmt"
    "os"
)

type Env struct {
    JwtSecret []byte
    UploadsUrl string
    // used to know when the environment has been loaded
    loaded bool
}

var env_keys = []string{"JWT_SECRET_KEY", "UPLOADS_URL"}

var env Env

func GetEnv() Env {
    if env.loaded {
        return env
    }

    for _, env_key := range env_keys {
        if _, exists := os.LookupEnv(env_key); !exists {
            fmt.Fprintf(os.Stderr, "[ERROR] \"%s\" variable is not declared on \".env\" file\n", env_key)
        }
    }

    env = Env{
        JwtSecret: []byte(os.Getenv("JWT_SECRET_KEY")),
        UploadsUrl: os.Getenv("UPLOADS_URL"),
        loaded: true,
    }

    return env
}
