//+build dev

package client

import (
    "net/http"
    "os"
)

var Client http.FileSystem = http.Dir("app"+ string(os.PathSeparator) +"client")
