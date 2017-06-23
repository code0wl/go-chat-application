package connector

import (
    "fmt"
    "net"
    "log"
    "bufio"
    "os"
)

// Channel takes an ip and port
// runs channel connection for a channel
func Channel(ip string, port string) {
    fmt.Println("channel on: " + ip + ":" + port)

    listener, err := net.Listen("tcp", ip+":"+port)
    if err != nil {
        log.Fatal("Oops: ", err)
    }

    conn, err := listener.Accept()
    if err != nil {
        log.Fatal("Oops: ", err)
    }
    fmt.Println("Channel has a new active user")

    for {
        startChannelEngine(conn)
    }
}

func startChannelEngine(conn net.Conn) {
    reader := bufio.NewReader(conn)
    message, readErr := reader.ReadString('\n')
    if readErr != nil {
        log.Fatal("Error: ", readErr)
    }
    fmt.Println("Message received: ", message)

    fmt.Print("Message: ")
    replyReader := bufio.NewReader(os.Stdin)
    replyMessage, replyErr := replyReader.ReadString('\n')
    if replyErr != nil {
        log.Fatal("Error: ", replyErr)
    }
    fmt.Fprint(conn, replyMessage)
}

func startSubscriberEngine(conn net.Conn) {
    fmt.Print("Send message: ")
    reader := bufio.NewReader(os.Stdin)
    message, readErr := reader.ReadString('\n')
    if readErr != nil {
        log.Fatal("Error: ", readErr)
    }
    fmt.Fprint(conn, message)

    replyReader := bufio.NewReader(conn)
    replyMessage, replyErr := replyReader.ReadString('\n')
    if replyErr != nil {
        log.Fatal("Error: ", replyErr)
    }
    fmt.Println("Message received:", replyMessage)
}

// Subscriber takes an ip and port
// runs subscriber connection to an existing channel
func Subscriber(ip string, port string) {
    fmt.Println("subscriber on: " + ip + ":" + port)
    conn, err := net.Dial("tcp", ip+":"+port)
    if err != nil {
        log.Fatal("Oops: ", err)
    }
    for {
        startSubscriberEngine(conn)
    }
}
