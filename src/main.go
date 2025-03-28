package main;

import (
    "net/http"
    "html/template"
    "path"
    "fmt"
);

const DEVELOPMENT = true;
const PORT = 3000;
const VIEWS_DIR = "./views";
const STATIC_DIR = "./static";
const BUILD_DIR = "./build";

type BlogPost struct {
    Id int;
    Title string;
    Cover string;
};

type HomePageData struct {
    Dev bool;
    BlogPosts []BlogPost;
};

func main() {
    staticFs := http.FileServer(http.Dir(STATIC_DIR));
    http.Handle("/static/", http.StripPrefix("/static/", staticFs));

    buildFs := http.FileServer(http.Dir(BUILD_DIR));
    http.Handle("/build/", http.StripPrefix("/build/", buildFs));

    homePath := path.Join(VIEWS_DIR, "home.html");
    basePath := path.Join(VIEWS_DIR, "layout/base.html");
    baseTmpl := template.Must(template.ParseFiles(basePath, homePath));

    http.HandleFunc("/",  func(w http.ResponseWriter, r *http.Request) {
        baseTmpl.Execute(w, HomePageData{
            Dev: DEVELOPMENT,
            BlogPosts: []BlogPost {
                { Id: 1, Title: "Test", Cover: "https://fvtblog.com/assets/covers/blog/8014-1732749325489.webp" },
            },
        });
    });

    fmt.Println("[INFO] Server listening on port", PORT);
    http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil);
}
