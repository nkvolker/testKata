package main

import ("fmt"; "bufio"; "os"; "strings")

func inputData() []string {
    scanner := bufio.NewScanner(os.Stdin)
    _ = scanner.Scan()
    input := strings.Split(scanner.Text(), " ")
    
    return input
}

func isValidInput(input []string) (x, y int, oper, format string) {
    switch {
        case len(input) != 3:
            panic("incorrect input data format")
        case arabic[input[0]] != 0 && arabic[input[2]] != 0:
            format = "arabic"; x = arabic[input[0]]; y = arabic[input[2]]
        case romans[input[0]] != 0 && romans[input[2]] != 0:
            format = "romans"; x = romans[input[0]]; y = romans[input[2]]
            if strings.Contains("-", input[1]) && x <= y {
                panic("no negative numbers and zero in the Roman system")
            }
        default: panic("incorrect input data")
    }
    if strings.Contains("+-*/", input[1]) {
        oper = input[1]
    } else {
        panic("string is not a mathematical operation")
    }
    return
}

func calculate(x, y int, oper string) (res int) {
    switch oper {
        case "+": res = x + y
        case "-": res = x - y
        case "*": res = x * y
        case "/": res = x / y
    }
    return
}

func output(res int, format string) {
    if format == "arabic" {
        fmt.Println(res)
    } else {
        var conv = map[int]string{
        100: "C", 90: "XC", 50: "L", 40: "XL", 10: "X", 9: "IX", 5: "V", 4: "IV", 1: "I"}
        convNumbers := [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
        
        for _, v := range convNumbers {
            for v <= res {
                fmt.Println(conv[v])
                res -= v
            }
        }
    }
}

var arabic = map[string]int{
    "1": 1, "2": 2, "3": 3, "4": 4, "5": 5,
    "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}
var romans = map[string]int{
    "I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
    "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

func main() {

    input := inputData()
    
    x, y, oper, format := isValidInput(input)
    
    res := calculate(x, y, oper)
    
    output(res, format)    
}
