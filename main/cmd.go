package main
import "fmt"

//определите функцию download()


func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

    var s1, s2, s3 int
    fmt.Scanln(&s1)
    fmt.Scanln(&s2)
    fmt.Scanln(&s3)

    go download(s1, ch1)
    go download(s2, ch2)
    go download(s3, ch3)

    func download(s, ch int) {
    for i:=0; i<=s; i++ {
       var sum int
        sum += i 
    }
    ch <- sum
}
}