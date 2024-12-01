package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "regexp"
    "strings"
    "sync"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run script.go /path/")
        return
    }

    root := os.Args[1]
    patterns := []string{
                "787backlink.xyz",
                "binusrun.com",
                "seokeras.com",
                "jasaiklan.lol",
                "umrohkemenag.org",
                "a.mikomallkopo.com",
                "purefine.online",
                "twins-4e372.web.app",
                "forbesstream.com",
                "pinjam100.site",
                "fixtechcs.net",
                "dw88.cwhonors.org",
                "w88.sepangracingteam.com",
                "ant1rungk4d",
                "ampshopify.store",
                "petirjago.site",
                "seoinclude.pages.dev",
                "includes-page.com",
                "wp.rpka.org",
                "admin.rpka.org",
                "rpka.org",
                "hedon77official.xyz",
                "linkgacorthailand.xyz",
                "dp288.live",
                "bobabotui.lol",
                "seobeton.team",
                "aHR0cHM6Ly95b2licmUuY29tL3NoZWxsL2FkbWluLnR4dA",
                "judi",
                "jackpot",
                "pragmatic",
                "casino",
                "winrate",
                "gacor",
                "togel",
                "maxwin",
                "maxbet",
                "scatter",
                "nextcloud",
                "owncloud",
        }

    regex := regexp.MustCompile(`(?i)` + strings.Join(patterns, "|"))

    var wg sync.WaitGroup
    fileChan := make(chan string)

    go func() {
        err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.IsDir() {
                fileChan <- path
            }
            return nil
        })
        if err != nil {
            fmt.Println("Error:", err)
        }
        close(fileChan)
    }()

    for i := 0; i < 10; i++ { // Number of goroutines
        wg.Add(1)
        go func() {
            defer wg.Done()
            for path := range fileChan {
                scanFile(path, regex)
            }
        }()
    }

    wg.Wait()
}

func scanFile(path string, regex *regexp.Regexp) {
    file, err := os.Open(path)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    lineNumber := 1
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Error reading file:", err)
            return
        }
        if regex.MatchString(line) {
            matchedPattern := regex.FindString(line)
            fmt.Printf("%s keyword %s at line %d\n", path, matchedPattern, lineNumber)
        }
        lineNumber++
    }
}
