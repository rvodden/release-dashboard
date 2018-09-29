package controller

import (
    "log"
    "net/http"
    "strings"
)

type neuteredFileSystem struct {
    fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
    log.Println("Opening " + path)

    f, err := nfs.fs.Open(path)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    s, err := f.Stat()
    if s.IsDir() {
        index := strings.TrimSuffix(path, "/") + "/index.html"
        if _, err := nfs.fs.Open(index); err != nil {
            log.Println(err)
            return nil, err
        }
    }

    return f, nil
}
