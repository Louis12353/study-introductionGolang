package main


func main() {

    testprinting()

    //1. declaring variables in golang
    // example of var declaration
    var x int = 5
    var y = "hello"
    
    // example of using shorthand := for local variables
    z := 42
    msg := "Hi, Go!"

    fmt.Println("x:", x)
    fmt.Println("y:", y)
    fmt.Println("z:", z)
    fmt.Println("msg:", msg)

    // 2. Functions: Declaring and using functions
    result := greet("Alice")
    fmt.Println("Greeting:", result)

    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient, "Remainder:", remainder)

    // 3. Loops: for loop, while style, infinite loop
    fmt.Println("\n--- Loops ---")
    for i := 0; i < 5; i++ {
        fmt.Println("For loop iteration", i)
    }

    x := 0
    for x < 3 {
        fmt.Println("While loop iteration", x)
        x++
    }

    fmt.Println("Infinite loop (break after 1 iteration):")
    for {
        fmt.Println("Looping forever")
        break
    }

    // 4. Conditionals: if, if with short statement, and switch
    fmt.Println("\n--- Conditionals ---")
    age := 20
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }

    // if with short statement
    if n := len("hello"); n > 3 {
        fmt.Println("Long word")
    }

    // Switch statement
    day := "Monday"
    switch day {
    case "Monday":
        fmt.Println("Start of week")
    case "Friday":
        fmt.Println("Weekend soon!")
    default:
        fmt.Println("Midweek")
    }

    // 5. Arrays, Slices, and Maps

    // Arrays (fixed size)
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println("\n--- Array ---")
    fmt.Println("Array elements:", arr)
    fmt.Println("Access first element:", arr[0])

    // Slices (dynamic size)
    nums := []int{10, 20, 30}
    nums = append(nums, 40)  // Add element
    fmt.Println("\n--- Slice ---")
    fmt.Println("Slice elements:", nums)
    fmt.Println("Slice from index 1 to 3:", nums[1:3])

    // Maps (key-value pairs)
    scores := map[string]int{
        "Alice": 90,
        "Bob":   85,
    }
    scores["Charlie"] = 95  // Adding new entry
    fmt.Println("\n--- Map ---")
    fmt.Println("Scores:", scores)
    fmt.Println("Bob's score:", scores["Bob"])
}

// Function Example 1: Greeting function
func greet(name string) string {
    return "Hello, " + name
}

// Function Example 2: Divide function that returns two values (quotient and remainder)
func divide(a, b int) (int, int) {
    return a / b, a % b
}
