package leetcode

// 哈希表 数组 map 

func twoSum(nums []int, target int) []int {

    m := make(map[int]int)

    for idx, num := range(nums) {
        if i, ok := m[target-num]; ok {
            return []int{i, idx}
        }
        m[num] = idx
    }

    return []int{}
}
