import "strconv"

func fizzBuzz(n int) []string {
	// Initialize an array with initial capacity
    result := make([]string, 0, n)
    for i := 0; i < n;i++ {
        value := i + 1
        if isDivisorOf3And5(value) {
            result = append(result, "FizzBuzz")
        } else if isDivisorOf3(value) {
            result = append(result, "Fizz")
        } else if isDivisorOf5(value) {
            result = append(result, "Buzz")
        } else {
            s := strconv.Itoa(value) // converting integer to string in go
            result = append(result, s)
        }
    }
    return result 
}

func isDivisorOf3And5(x int) bool {
    if x % 3 == 0 && x % 5 == 0 {
        return true
    }
    return false
}

func isDivisorOf3(x int) bool {
    if x % 3 == 0 {
        return true
    }
    return false
}

func isDivisorOf5(x int) bool {
    if x % 5 == 0 {
        return true
    }
    return false
}