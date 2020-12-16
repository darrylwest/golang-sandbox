package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

const (
    RuneStart = rune('a')
    RuneEnd = rune('z')
    RuneCount = RuneEnd - RuneStart + 1
)

var (
    debug = false
)

type TrieType interface {
    Find(name string) *Branch
    Add(name string) *Branch
}

type Branch struct {
    value rune
    children map[rune]*Branch
    isWord bool
    count int
}

type Trie struct {
    branches [RuneCount]Branch
}

func (t *Trie) Find(word string) int {

    rs := []rune(word)
    idx := rs[0] - RuneStart
    root := t.branches[idx]
    if root.value != rs[0] {
        return 0
    }

    if debug { fmt.Printf("found root for %c\n", root.value) }
    var branch = &root

    for i := 1; i < len(rs); i++ {
        r := rs[i]
        next := branch.findChild(r)
        if next == nil {
            return 0
        }

        if debug { fmt.Printf("found %c %v\n", r, next.count) }

        branch = next
    }

    // if we made it this far, then we have partials so find and 
    // return all the full words under this branch

    if debug { fmt.Printf("found %s at branch %v\n", word, branch) }

    return branch.count
}

// word (string) to []rune and []rune to string
func (t *Trie) Add(word string) Branch {
    rs := []rune(word);

    idx := rs[0] - RuneStart
    root := t.branches[ idx ]
    if root.value != rs[0] {
        root.value = rs[0]
        root.children = map[rune]*Branch{}
    }

    root.count++
    t.branches[ idx ] = root

    var branch *Branch = &root

    for i := 1; i < len(rs); i++ {
        r := rs[i]
        next := branch.findChild(r)

        if next == nil {
            child := Branch{ value: r, children: map[rune]*Branch{}, count:1 }
            branch.children[r] = &child
            if debug { fmt.Printf("add %v %d\n", child, len(branch.children)); }
            branch = &child
        } else {
            if branch != &root {
                branch.count++
            }
            if debug { fmt.Printf("found %v %d\n", branch, len(branch.children)); }
            branch = next
        }
    }

    branch.isWord = true

    if debug { fmt.Printf("Add root: %v, branch %v\n", root, branch) }

    return root
}

func (b *Branch) findChild(r rune) *Branch {
    if child, ok := b.children[r]; ok {
        return child
    }

    return nil
}

func (t Trie) Dump() {
    fmt.Println("branch count:", len(t.branches));
    for _,branch := range t.branches {
        if branch.value >= RuneStart {
            fmt.Printf("%c -> %v\n", branch.value, len(branch.children))
            branch.Dump()
        }
    }
}

func (b *Branch) Dump() {
    for k,v := range b.children {
        fmt.Printf("%c -> %c %v %d\n", b.value, k, v, len(v.children))
        if len(v.children) > 0 {
            v.Dump()
        }
    }
}

func main() {
    trie := new(Trie)

    if false {
        words := readDictionary("words.txt")

        indexes := []int{ 200, 17130, 28197, 28198, 48100, 48116 }
        
        for _,i := range indexes {
            trie.Add(words[i]);
        }

        word := words[48100]
        count := trie.Find(word)
        fmt.Printf("word %s count %d\n", word, count)

        word = words[48116]
        count = trie.Find(word)
        fmt.Printf("word %s count %d\n", word, count)
    } else if false {
        trie.Add("hack")
        trie.Add("hackerrank")

        // now dump the entire trie for visual inspection...
        if debug {
            trie.Dump()
        }

        word := "hac"
        count := trie.Find(word)
        fmt.Printf("word %s count %d\n", word, count)

        word = "hak"
        count = trie.Find(word)
        fmt.Printf("word %s count %d\n", word, count)
    } else {
        // debug = true
        words := []string{ "s", "ss", "sss", "ssss", "sssss" }
        for _, word := range words {
            trie.Add(word)
        }
        expect := []int{ 5, 4, 3, 2, 1 }
        // for i, v := range expect[1:2] {
            v := expect[1]
            debug = true
            word := words[1]
            count := trie.Find(word)
            fmt.Printf("search for %s found %d expect %d\n", word, count, v)
        //}
        // trie.Dump()
    }

}

func readDictionary(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }

    defer file.Close()
    words := []string{}

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := strings.ToLower(scanner.Text())
        words = append(words, word)
    }

    fmt.Println(len(words), "words loaded...");

    return words
}
