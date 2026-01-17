package main

import (
    "flag"
    "fmt"
    "os"
    "regexp"
    "sort"
)

type wordCount struct {
    word  string
    count int
}

func main() {
    top := flag.Int("top", 10, "top N words")
    flag.Parse()
    if flag.NArg() < 1 {
        fmt.Fprintln(os.Stderr, "usage: ./textstats <path> [--top N]")
        os.Exit(1)
    }
    path := flag.Arg(0)
    data, err := os.ReadFile(path)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    text := string(data)

    lines := 0
    if len(text) > 0 {
        lines = 1
        for _, ch := range text {
            if ch == '\n' {
                lines++
            }
        }
    }

    re := regexp.MustCompile("[A-Za-z0-9]+")
    words := re.FindAllString(text, -1)
    counts := map[string]int{}
    for _, w := range words {
        lower := []rune(w)
        for i, ch := range lower {
            if ch >= 'A' && ch <= 'Z' {
                lower[i] = ch + 32
            }
        }
        key := string(lower)
        counts[key]++
    }

    list := make([]wordCount, 0, len(counts))
    for w, c := range counts {
        list = append(list, wordCount{word: w, count: c})
    }
    sort.Slice(list, func(i, j int) bool {
        if list[i].count == list[j].count {
            return list[i].word < list[j].word
        }
        return list[i].count > list[j].count
    })

    fmt.Printf("lines: %d\n", lines)
    fmt.Printf("words: %d\n", len(words))
    fmt.Printf("chars: %d\n", len(text))
    fmt.Println("top words:")
    limit := *top
    if limit > len(list) {
        limit = len(list)
    }
    for i := 0; i < limit; i++ {
        fmt.Printf("  %s: %d\n", list[i].word, list[i].count)
    }
}
