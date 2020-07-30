package main

import (
    "fmt"
)

type TrieNode  struct {
    children   []*TrieNode
    isWord      bool
    val         string
    count       int
}


func (root *TrieNode) Insert (word string) {
    cur := root
    for i := range word {
        c := int(word[i]-'a')
        if cur.children[c] == nil {
            cur.children[c] = &TrieNode{make([]*TrieNode, 26), false, string(c), 0}
        }
        cur.count++
        cur = cur.children[c]
    }
    cur.count++
    cur.isWord = true
}

func (root *TrieNode) Search (word string) bool {
    cur := root
    for i := range word {
        c := int(word[i]-'a')
        if cur.children[c] == nil {
            return false
        }
        cur = cur.children[c]
    }
    return cur.isWord
}

func (root *TrieNode) CountPerfix (perfix string) int {
    cur := root
    for i := range perfix {
        c := int(perfix[i]-'a')
        if cur.children[c] == nil {
            return 0
        }
        cur = cur.children[c]
    }
    return cur.count
}

func main() {
    root := &TrieNode{make([]*TrieNode, 26), false, "", 0}
    root.Insert("foo")
    root.Insert("foon")
    root.Insert("foaaa")
    fmt.Println("words count is", root.count)
    fmt.Println("Search foon", root.Search("foon"))
    fmt.Println("Search foo", root.Search("foo"))
    fmt.Println("Search bar", root.Search("bar"))
    fmt.Println("Perfix count foo is ", root.CountPerfix("foo"))
    fmt.Println("Perfix count fo is ", root.CountPerfix("fo"))

}
