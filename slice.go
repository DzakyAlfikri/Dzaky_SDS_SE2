
package main

import "fmt"

func main() {
    data := []int{1, 2, 3, 4, 5}
    fmt.Println("data awal:", data)
    data = append(data, 6, 7, 8)

    fmt.Println("Setelah penambahan:", data)
}
