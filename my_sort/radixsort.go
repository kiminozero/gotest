package mysort



func LSDSort(nums []int) {
    if len(nums) < 2 {
        return
    }
    maxVal := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > maxVal {
            maxVal = nums[i]
        }
    }

    exp := 1
    radix := 10

    aux := make([]int, len(nums))
    for maxVal / exp > 0 {
        count := make([]int, radix)
        // Counting sort
        for i := 0; i < len(nums); i++ {
            count[(nums[i]/exp)%10]++
        }

        // you could also use partial_sum()
        for i := 1; i < len(count); i++ {
            count[i] += count[i-1]
        }
        // 
        for i := len(nums)-1; i > -1; i-- {
            count[(nums[i]/exp)%10]--
            aux[count[(nums[i]/exp)%10]] = nums[i]
        }

        copy(nums, aux)
        exp *= 10
    }
}

