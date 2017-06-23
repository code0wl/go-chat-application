package main

import (
    "flag"
    "fmt"
)

// package member vars
var (
    isMainChannel bool
    listen = string("Listens to IP and port")
)

func main() {

    flag.BoolVar(&isMainChannel, "listen", false, listen)
    flag.Parse()

    if isMainChannel {
        fmt.Print("channel")
    } else {
        fmt.Print("subscriber")
    }

}
