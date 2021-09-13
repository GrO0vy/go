package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 난수 추출된 수의 소수 판정 프로그램 ver 0.2
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0 과 1 제외
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 { // 1과 number 일 때 loop를 돌지 않음
			count = count + 1
		}
	}
	if count == 2 {
		fmt.Printf("%d는 소수입니다.", number)
	} else {
		fmt.Printf("%d는 소수가 아닙니다.", number)
	}

}
