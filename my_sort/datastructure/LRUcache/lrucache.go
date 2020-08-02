package main

import(
    "fmt"
)

type LRUCache struct {
    head, tail *node
    cache       map[int]*node
    capacity    int
    size        int
}

type node struct {
    key     int
    value   int
    prev    *node
    next    *node
}
func (this *LRUCache) removeTail() {
    prev := this.tail.prev
    prev.prev.next = this.tail
    this.tail.prev = prev.prev
    delete(this.cache, prev.key)
}
func (this *LRUCache) remove(key int) {
    prev := this.cache[key].prev
    prev.next = this.cache[key].next
    prev.next.prev = prev
    delete(this.cache, key)
}

func (this *LRUCache) add(key, value int) {
    this.cache[key] = &node{key, value, nil, nil}
    next := this.head.next
    next.prev = this.cache[key]
    this.head.next = this.cache[key]
    this.cache[key].next = next
    this.cache[key].prev = this.head
}

func (this *LRUCache) goHead(key, value int) {
    this.remove(key)
    this.add(key, value)
}

func Constructor(capacity int) LRUCache {
    head := &node{}
    tail := &node{}
    head.next = tail
    tail.prev = head
    return LRUCache{head, tail, make(map[int]*node), capacity, 0}
}


func (this *LRUCache) Get(key int) int {
    // - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
    if _, ok := this.cache[key]; ok {
        v := this.cache[key].value
        this.goHead(key, v)
        return v
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.cache[key]; !ok {
        // 如果密钥不存在，则写入其数据值
        if this.size < this.capacity {
            this.size++
        }else {
            this.removeTail()
        }
        this.add(key, value)
    } else {
        // 存在 更新值 把节点挪到head
        this.goHead(key, value)
    }

}




/**
OBOBOB * Your LRUCache object will be instantiated and called as such:
OBOBOB * obj := Constructor(capacity);
OBOBOB * param_1 := obj.Get(key);
 * obj.Put(key,value);
OBOBOB */


func main() {

    cache := Constructor(2)
    fmt.Println("Put(1,1)")
    cache.Put(1, 1)
    cache.Put(2, 2)
    fmt.Println(cache.Get(1));       // 返回  1
    cache.Put(3, 3)                  // 该操作会使得关键字 2 作废
    fmt.Println(cache.Get(2));       // 返回 -1 (未找到)
    cache.Put(4, 4)                 // 该操作会使得关键字 1 作废
    fmt.Println(cache.Get(1));       // 返回 -1 (未找到)
    fmt.Println(cache.Get(3));       // 返回  3
    fmt.Println(cache.Get(4));       // 返回  4

}
