package main

import (
    "net/http"
    "fmt"
    "log"
    "io/ioutil"
    "os"
    "github.com/sergiubcn/golang-demo/handlers"
)
func main() {
    l := log.New(os.Stdout, "product-api", log.LstdFlags)
    hh := handlers.NewHello(l)

    sm := http.NewServeMux()
    sm.Handle("/", hh)

    http.ListenAndServe(":9090", nil)
}