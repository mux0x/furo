package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strings"
)

func main() {
    inputFile := flag.String("i", "", "file that contains output")
    outputFile := flag.String("o", "nonFFUFdups.txt", "location of output file")
    flag.Parse()

    if *inputFile == "" {
        fmt.Println("[-] you need to provide a file that contains urls e.g., -i file.txt")
        os.Exit(0)
    }

    writtenValues := make(map[string]bool)

    file, err := os.Open(*inputFile)
    if err != nil {
        fmt.Printf("[-] error opening file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        x := strings.TrimSpace(line)
        length := strings.Split(x, ",")[1]
        if _, ok := writtenValues[length]; !ok {
            writtenValues[length] = true
            fmt.Println(x)
            f, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
            if err != nil {
                fmt.Printf("[-] error opening file: %v\n", err)
                os.Exit(1)
            }
            defer f.Close()

            fmt.Fprintln(f, x)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("[-] error scanning file: %v\n", err)
        os.Exit(1)
    }
}

