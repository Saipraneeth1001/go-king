	// 1st approach
func singleNumberOne(nums []int) int {

    // Initialize a map to maintain the counts
    var countMap map[int]int
    countMap = make(map[int]int)

    for _, value := range nums {
        if _, ok := countMap[value]; ok {
            countMap[value] = countMap[value] + 1
        } else {
            countMap[value] = 1
        }
    }

    for key, value := range countMap {
        if value == 1 {
            return key
        }
    }
    return -1
}

// The fastest and xor approach

func singleNumber(nums []int) int {
    var result int
    for _, num := range nums {
        result = result ^ num
    }
    return result
}

/**
| Property        | Explanation                            | Example                     |
| --------------- | -------------------------------------- | --------------------------- |
| `x ^ x = 0`     | A number XORed with itself cancels out | `4 ^ 4 = 0`                 |
| `x ^ 0 = x`     | A number XORed with 0 remains the same | `4 ^ 0 = 4`                 |
| **Commutative** | Order doesn’t matter                   | `a ^ b ^ c = c ^ a ^ b`     |
| **Associative** | Grouping doesn’t matter                | `(a ^ b) ^ c = a ^ (b ^ c)` |

*/