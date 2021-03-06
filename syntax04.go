package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

// 난수 추출된 수의 소수 판정 프로그램 v0.6
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true // 소수인가?
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break //2부터해서 약수가 발견되면 반복문 즉시 종료
		}
		fmt.Print(i, " ")
	}

	if isPrime {
		fmt.Println(number, "는(은) 소수입니다")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~")
	}
}
