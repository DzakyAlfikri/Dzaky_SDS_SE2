package main

import "fmt"

func main() {
    data := []int{16, 11, 7, 9, 32, 1, 26, 5, 31, 11, 3}

    bubbleSort(data)

    fmt.Println("Data terurut dari terbesar ke terkecil:", data)
}

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            // Jika elemen saat ini lebih kecil dari elemen berikutnya, tukar
            if arr[j] < arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}
