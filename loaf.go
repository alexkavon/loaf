package main

import (
    "fmt"
    "gitlab.com/alexkavon/loaf/models"
    "gitlab.com/alexkavon/loaf/store"
)

func main() {
    s, err := store.NewStore()
    if err != nil {
        fmt.Println(err)
    }
    loaf := models.NewLoaf("alex")
    fmt.Println(s)
    fmt.Println(loaf)
}
