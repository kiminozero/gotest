package main

import(
    "fmt"
)


type Trie struct {
        num  int //有多少单词通过这个节点,即由根至该节点组成的字符串模式出现的次数
        son  []*Trie //所有的儿子节点
        isEnd  bool //是不是最后一个节点
        val    byte //节点的值
        haveSon  bool
}
const SIZE = 26

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        num : 1,
        son : make([]*Trie, SIZE),
        // isEnd : false,
        // haveSon : false,
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
        if word==""||len(word)==0 {
            return
        }
        node :=this;
        for i := 0; i<len(word); i++ {
            pos :=word[i]-'a';
            // fmt.Println(pos, node.son[pos])
            if node.son[pos] == nil {
                node.haveSon = true;
                node.son[pos] = &Trie{
                     val : word[i],
                     num: 1,
                     son : make([]*Trie, SIZE),
                     };
            } else {
                node.son[pos].num++;
            }
            node=node.son[pos];
        }
        node.isEnd=true;
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    if word=="" || len(word)==0 {
                return false
    }
    node :=this;
    for i:=0; i<len(word); i++ {
        pos :=word[i]-'a'
        if(node.son[pos]!=nil) {
            node=node.son[pos]
        } else {
            return false
        }
    }
    return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    if  prefix == "" || len(prefix) == 0 {
            return true
        }
    node := this

    for i := 0; i < len(prefix); i++ {
        pos := prefix[i] - 'a';
        if (node.son[pos] == nil) {
            return false
        } else {
            node = node.son[pos];
        }
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

func main() {
    obj := Constructor()
    obj.Insert("foo")
    obj.Insert("fooo")
    obj.Insert("fooaa")
    obj.Insert("bar")
    fmt.Println(obj.Search("foo"))
    fmt.Println(obj.StartsWith("foo"))
}
