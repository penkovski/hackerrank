package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    nums := make(map[int]int)
    
    for i := 0; i < n; i++ {
        var f int
        fmt.Scan(&f)
        nums[f] = f
    }
    
    for i := 1; i <= n; i++ {
        _, ok := nums[i]
        if !ok {
            fmt.Println("NO")
            return
        }
    }
    
    fmt.Println("YES")
}

