package main

func main() {
    v, err := divide(10, 0)
    if err != nil {
      fmt.Println("uh oh:", err)
    }
    fmt.Println("hello, %f", v)
}

func divide(n, d float64) (float64, error) {
    return n / d   
}
