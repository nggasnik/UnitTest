package main

import (
	"errors"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type Wallet struct {
	PersonName string
	Credit     float64
	mutex      sync.Mutex
}

// Put your logic and code inside this function.
func (w *Wallet) Withdrawal(wdAmount float64) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if wdAmount > w.Credit {
		return errors.New("insufficient balance")
	}
	w.Credit -= wdAmount
	return nil
}

// Put your logic and code inside this function.
func (w *Wallet) GetWallet() Wallet {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return *w
}

// ===== DO NOT EDIT. =====
const minWd = 1
const maxWd = 20

var wallet = &Wallet{
	PersonName: "John Doe",
	Credit:     maxWd,
	mutex:      sync.Mutex{},
}

func TestCaseWallet(t *testing.T) {
	iteration := 20
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	successfulWithdrawal := false

	for i := 0; i <= iteration; i++ {
		wdAmount := float64(rand.Intn((maxWd-minWd)/0.05))*0.05 + minWd

		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			remaining, err := atm(x, wdAmount, wallet)
			if err == nil {
				log.Printf("Withdraw Amount: %.2f, Remaining Credit: %.2f", wdAmount, remaining)
				successfulWithdrawal = true
			}
		}(i)
	}

	wg.Wait()
	log.Println("+------------+")
	log.Printf("%s's final credit: %.2f", wallet.PersonName, wallet.Credit)

	switch {
	case wallet.Credit < 0:
		t.Fail()
	case !successfulWithdrawal:
		t.Fail()
	}
}

func atm(
	c int,
	wdAmount float64,
	wal *Wallet,
) (float64, error) {
	err := wal.Withdrawal(wdAmount)
	if err != nil {
		return wal.GetWallet().Credit, err
	}

	return wal.GetWallet().Credit, nil
}

func main() {
	testSuite := []testing.InternalTest{
		{
			Name: "TestCaseWallet",
			F:    TestCaseWallet,
		},
	}

	testing.Main(nil, testSuite, nil, nil)
}

// ===== DO NOT EDIT. =====


// ======= HASIL OUTPUT =====

// MacBook-Pro TEST 2 % go run main.go  
// 2023/02/23 14:21:04 Withdraw Amount: 11.70, Remaining Credit: 8.30
// 2023/02/23 14:21:04 Withdraw Amount: 8.20, Remaining Credit: 0.10
// 2023/02/23 14:21:04 +------------+
// 2023/02/23 14:21:04 John Doe's final credit: 0.10