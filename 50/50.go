package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 1; i <= 50; i++ {
		rand.Seed(time.Now().UnixNano())
		a := rand.Intn(45) + 5
		time.Sleep(10 * time.Millisecond)
		b := 0
		b = b + 1
		if a <= 15 {
			fmt.Printf("%d 번째 손님 소요시간 %d 분입니다.\n", b, a)
		}
	}
}
