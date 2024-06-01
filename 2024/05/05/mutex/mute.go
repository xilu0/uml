package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 1000 // 初始余额
}

func deposit(value int, wg *sync.WaitGroup) {
	mutex.Lock() // 请求互斥锁
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock() // 释放互斥锁
	wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock() // 请求互斥锁
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
	balance -= value
	mutex.Unlock() // 释放互斥锁
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(500, &wg)
	go withdraw(700, &wg)
	wg.Wait()

	fmt.Printf("New Balance: %d\n", balance)
}
