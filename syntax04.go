package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 난수 추출된 수의 소수 판정 프로그램 ver 0.4
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2 // 0 과 1 제외
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 { // 1과 number 일 때 loop를 돌지 않음
			isPrime = false
		}
		fmt.Print(i, " ")
	}
	if isPrime { // 비교연산 제거
		fmt.Printf("%d는 소수입니다.", number)
	} else {
		fmt.Printf("%d는 소수가 아닙니다.", number)
	}

}
