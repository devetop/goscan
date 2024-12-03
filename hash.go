package main

import (
    "crypto/sha256"
    "encoding/hex"
    "flag"
    "fmt"
    "io"
    "log"
    "os"
    "path/filepath"
)

func main() {
    pathPtr := flag.String("path", ".", "Path to directory")
    hashPtr := flag.String("hash", "", "SHA256 hash to match")
    flag.Parse()

    if *hashPtr == "" {
        log.Fatal("Hash value must be provided")
    }

    err := filepath.Walk(*pathPtr, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            fileHash, err := hashFileSha256(path)
            if err != nil {
                log.Println("Error hashing file:", err)
                return nil
            }

            if fileHash == *hashPtr {
                fmt.Printf("File matched: %s\n", path)
            }
        }
        return nil
    })

    if err != nil {
        log.Println("Error walking the path:", err)
    }
}

func hashFileSha256(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    hash := sha256.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }

    return hex.EncodeToString(hash.Sum(nil)), nil
}
