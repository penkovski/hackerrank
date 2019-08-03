package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)

func main() {
    var x1, v1 int64
    var x2, v2 int64
    
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        s := scanner.Text()
        args := strings.Fields(s)
        
        x1, _ = strconv.ParseInt(args[0], 10, 64)
        v1, _ = strconv.ParseInt(args[1], 10, 64)
        x2, _ = strconv.ParseInt(args[2], 10, 64)
        v2, _ = strconv.ParseInt(args[3], 10, 64)        
    }
    
    if x2 > x1 && v2 > v1 {
        fmt.Println("NO")
        return
    }
    
    if x1 > x2 && v1 > v2 {
        fmt.Println("NO")
        return
    }
    
    if v2 - v1 == 0 {
        fmt.Println("NO")
        return
    }
    
    if (x1 - x2) % (v2 - v1) == 0 {
        fmt.Println("YES")
        return
    }
    
    fmt.Println("NO")
}