package main
import "fmt"

func main() {
    var input string
    fmt.Scanln(&input)
    
    p := []rune("SOS")
    count := 0
    for i, r := range input {
        tmp := i % 3
        if p[tmp] != r {
            count += 1
        }
    }
    
    fmt.Println(count)
}

