# Simple radio functionality with Golang

I am new to go, just a simple program that uses tcp to allow communication in the same way a radio channel would work. 

###limitations
Message has to be received before one can be returned. 
Next step is to make it work concurrently.

# How to

Create a channel

```bash
go run main/go -listen "your port"
```

Subscribe to a channel

```bash
go run main/go "your port"
```

