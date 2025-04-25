package middlewares

import (
    "net/http"
    "github.com/FernandoVT10/go-blog/internals/controllers"
    "github.com/FernandoVT10/go-blog/internals/router"

    httpUtils "github.com/FernandoVT10/go-blog/internals/utils/http"
)

func AuthPage() router.Middleware {
    return func(w http.ResponseWriter, r *http.Request, next router.NextFunction) {
        if !controllers.IsAuthenticated(r) {
            http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
            return
        }
        next()
    }
}

func AuthApi() router.Middleware {
    return func(w http.ResponseWriter, r *http.Request, next router.NextFunction) {
        if !controllers.IsAuthenticated(r) {
            error := map[string]string{"error": "You don't have the right permissions"}
            httpUtils.SendJson(w, http.StatusUnauthorized, error)
            return
        }
        next()
    }
}
