package main

import (
    "fmt"
    "math/rand"
    "time"
)

// 1-------------->5-------------->9-------->11->nil
// 1------>3------>5------>7------>9-------->11->nil
// 1-->2-->3-->4-->5-->6-->7-->8-->9-->10--->11->nil


const (
    MaxLevel = 32
    Probability = 0.25
)

func randLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32() < Probability && level < MaxLevel; level++ {
	}
	return
}


type node struct {
    key     int
    next    []*node
}

type Skiplist struct {
    head    *node
    level   int
}


func Constructor() Skiplist {
    return Skiplist{&node{0, make([]*node, MaxLevel)}, 1}
}


func (this *Skiplist) Search(target int) bool {
    cur := this.head
    for i := this.level-1; i > -1; i-- {
        for cur.next[i] != nil {
            if cur.next[i].key == target {
                return true
            } else if cur.next[i].key > target {
                break
            } else {
                cur = cur.next[i]
            }
        }
    }
    return false
}


func (this *Skiplist) Add(num int)  {
    current := this.head
	update := make([]*node, MaxLevel) // 新节点插入以后的前驱节点
	for i := this.level - 1; i >= 0; i-- {
		if current.next[i] == nil || current.next[i].key > num {
			update[i] = current
		} else {
			for current.next[i] != nil && current.next[i].key < num {
				current = current.next[i] // 指针往前推进
			}
			update[i] = current
		}
	}

	level := randLevel()
	if level > this.level {
		// 新节点层数大于跳表当前层数时候, 现有层数 + 1 的 head 指向新节点
		for i := this.level; i < level; i++ {
			update[i] = this.head
		}
		this.level = level
	}
	node := &node{num, make([]*node, level)}
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}

}


func (this *Skiplist) Erase(num int) bool {
    
    flag := false
    cur := this.head
    for i := this.level-1; i > -1; i-- {
        for cur.next[i] != nil {
            if cur.next[i].key == num {
                cur.next[i], cur.next[i].next[i] = cur.next[i].next[i], nil
                flag = true
                break
            } else if cur.next[i].key > num {
                break
            } else {
                cur = cur.next[i]
            }
        }
    }

    return flag
}


/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

func main() {
    obj := Constructor();
    obj.Add(1);
    obj.Add(2);
    obj.Add(3);
    fmt.Println(obj.Search(4));
    obj.Add(4);
    fmt.Println("Search 1", obj.Search(1));
    fmt.Println("Erase 0", obj.Erase(0));
    fmt.Println("Erase 1", obj.Erase(1));
    fmt.Println("Search 1", obj.Search(1));
}
