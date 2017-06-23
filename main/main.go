package main

import (
    "flag"
    "os"
    "github.com/code0wl/go-chat/main/connector"
)

const (
    listen = string("Listens to IP and port")
    port   = string("8080")
)

func main() {
    var isMainChannel bool
    flag.BoolVar(&isMainChannel, "listen", false, listen)
    flag.Parse()

    if isMainChannel {
        ip := os.Args[2]
        connector.Channel(ip, port)
    } else {
        ip := os.Args[1]
        connector.Subscriber(ip, port)
    }
}
