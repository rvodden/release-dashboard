// +build ignore

package main

import (
    "awesomeProject/app/client"
    "github.com/shurcooL/vfsgen"
    "log"
    "os"
)

func main() {
    err := vfsgen.Generate(client.Client, vfsgen.Options{
        Filename: "app" + string(os.PathSeparator) + "client" + string(os.PathSeparator) +  "generated_client.go",
        PackageName:  "client",
        BuildTags:    "!dev",
        VariableName: "Client",
    })
    if err != nil {
        log.Fatalln(err)
    }
}
