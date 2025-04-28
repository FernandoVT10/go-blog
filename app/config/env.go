package config

import (
    "fmt"
    "os"
)

type Env struct {
    JwtSecret []byte
    UploadsUrl string
    Production bool
    MongoUri string
    MongoDbName string
    AdminPass string
    // used to know when the environment has been loaded
    loaded bool
}

var env_keys = []string{"JWT_SECRET_KEY", "UPLOADS_URL", "MONGO_URI", "MONGO_DB_NAME", "ADMIN_PASSWORD"}

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
        Production: os.Getenv("GO_ENV") == "production",
        MongoUri: os.Getenv("MONGO_URI"),
        MongoDbName: os.Getenv("MONGO_DB_NAME"),
        AdminPass: os.Getenv("ADMIN_PASSWORD"),
        loaded: true,
    }

    return env
}
