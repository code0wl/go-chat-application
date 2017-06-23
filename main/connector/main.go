package connector

import (
    "fmt"
    "net"
    "log"
    "bufio"
    "os"
)

func main() {

}

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

    reader := bufio.NewReader(conn)
    message, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal("Oops: ", err)
    }
    fmt.Println("Message send: ", message)
}

// Channel takes an ip and port
// runs subscriber connection to an existing channel
func Subscriber(ip string, port string) {
    fmt.Println("subscriber on: " + ip + ":" + port)
    conn, err := net.Dial("tcp", ip+":"+port)
    if err != nil {
        log.Fatal("Oops: ", err)
    }
    fmt.Print("Message :")

    reader := bufio.NewReader(os.Stdin)
    message, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal("Oops: ", err)
    }
    fmt.Fprint(conn, message)
}
