// slice return version

func plusOne(digits []int) []int {
    result := []int{}
    return compute(digits, result, 1, len(digits)-1)
}

func compute(nums []int, result []int, carry int, idx int) []int {
    if idx < 0 {
        if carry == 1 {
            return append([]int{1}, result...)
        }
        return result
    }

    addition := nums[idx] + carry

    if addition == 10 {
        result = append([]int{0}, result...)
        return compute(nums, result, 1, idx-1)
    } else {
        result = append([]int{addition}, result...)
        return compute(nums, result, 0, idx-1)
    }
}


// pointer version
func plusOnePointer(digits []int) []int {
    result := []int{}
    compute(digits, &result, 1, len(digits) - 1)
    return result
}

func computePointer(nums []int, result *[]int, carry int, currIndex int) {
    var addition int
    if currIndex < 0 && carry == 1 {
        *result = append([]int{1}, *result...)
        return
    } else if currIndex < 0 && carry == 0 {
        return
    }
    if carry == 1 {
        addition = nums[currIndex] + 1
    } else {
        addition = nums[currIndex]
    }

    if addition == 10 {
        *result = append([]int{0}, *result...)
        compute(nums, result, 1, currIndex - 1)
    } else {
        *result = append([]int{addition}, *result...)
        compute(nums, result, 0, currIndex - 1)
    }
}

// slice return version

