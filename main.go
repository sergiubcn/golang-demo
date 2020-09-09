package main

import (
    "fmt"
    "handlers/handlers"
    "net/http"
    // "io/ioutil"
)
func main() {
    test := "srg"
    users := handlers.NewUsers(test)
    usersHandler = handlers.NewUsers()

    sm := http.NewServeMux()
    sm.handle("/users", usersHandler)

    http.ListenAndServe(":9090", sm)
}