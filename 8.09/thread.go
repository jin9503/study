package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct { //은행계좌
	balance int //잔액 (메모리 보호대상) Mutex
	mutex   *sync.Mutex
}

func (a *Account) Widthdraw(val int) { //인출
	a.mutex.Lock()
	a.balance -= val
	a.mutex.Unlock()
}

func (a *Account) Deposit(val int) { //입금
	a.mutex.Lock()
	a.balance += val
	a.mutex.Unlock()
}

func (a *Account) Balance() int {
	a.mutex.Lock()
	balance := a.balance
	a.mutex.Unlock()
	return balance
}

var accounts []*Account //slice 배열 (메모리 보호대상) Mutex

func Transfer(sender, receiver int, money int) { //송금자, 받는자, 송금액
	accounts[sender].Widthdraw(money) // 동일하게 빼고
	accounts[receiver].Deposit(money) //동일하게 넣는데도 왜 잔액이 바뀌는가? 메모리가 엉클어 졌기 때문이다.
}

func GetTotalBalance() int { //전체 잔액량
	total := 0
	for i := 0; i < len(accounts); i++ {
		total += accounts[i].Balance()
	}
	return total
}

func RandomTransfer() { //랜덤함수 레퍼런스
	var sender, balance int //sender가 잔액이 있어야 하므로
	for {
		sender := rand.Intn(len(accounts))
		balance = accounts[sender].Balance()
		if balance > 0 { //샌더 잔액여부 확인
			break //샌더 잔액이 있으면 멈춘다
		}
	}

	var receiver int
	for {
		receiver = rand.Intn(len(accounts)) //Intn은 마지막값은 포함하지 않는 랜덤함수
		if sender != receiver {             //sender와 receiver가 다를때 까지
			break
		}
	}

	money := rand.Intn(balance) //송금량
	Transfer(sender, receiver, money)
}

func GoTransfer() { // 아래 쓰레드가 무한루프를 돌면서 랜덤하게 sender를뽑고 랜덤하게 잔액을 입출금 전송한다
	for {
		RandomTransfer()
	}
}

func PrintTotalBalance() {
	fmt.Printf("Total: %d\n", GetTotalBalance())
}

func main() {
	for i := 0; i < 20; i++ { //20개 어카운트 배열 1000 지정
		accounts = append(accounts, &Account{balance: 1000, mutex: &sync.Mutex{}}) //Mutex{}
	}

	PrintTotalBalance()

	for i := 0; i < 1; i++ {
		go GoTransfer() // for문을 10이하까지 루프로  GoTransfer함수를 수행하는 쓰레드를 10개를 만든다.
	}

	for {
		PrintTotalBalance()                // 해당부분도 쓰레드로 작동된다
		time.Sleep(100 * time.Millisecond) // 해당부분도 쓰레드로 작동
	}
}
