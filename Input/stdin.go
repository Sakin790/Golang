package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func main() {

    fmt.Print("Enter your name: ")

    r := bufio.NewReader(os.Stdin)

    name, err := r.ReadString('\n')

    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Hello %s!\n", strings.TrimSpace(name))
}