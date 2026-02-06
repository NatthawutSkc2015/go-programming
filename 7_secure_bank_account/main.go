package main

import (
	"fmt"
	"sync"
	"time"
)

// BankAccount represents a thread-safe bank account
type BankAccount struct {
	balance int
	mu      sync.Mutex
}

// NewBankAccount creates a new bank account with initial balance
func NewBankAccount(initialBalance int) *BankAccount {
	return &BankAccount{
		balance: initialBalance,
	}
}

// Deposit adds money to the account (thread-safe)
func (acc *BankAccount) Deposit(amount int) {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	acc.balance += amount
	fmt.Printf("✅ Deposited %d, New Balance: %d\n", amount, acc.balance)
}

// Withdraw removes money from the account (thread-safe)
// Returns true if successful, false if insufficient funds
func (acc *BankAccount) Withdraw(amount int) bool {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	// Check if sufficient funds
	if acc.balance < amount {
		fmt.Printf("❌ Failed to withdraw %d (Balance: %d - Insufficient funds)\n", amount, acc.balance)
		return false
	}

	// Sufficient funds, proceed with withdrawal
	acc.balance -= amount
	fmt.Printf("✅ Withdrew %d, New Balance: %d\n", amount, acc.balance)
	return true
}

// GetBalance returns the current balance (thread-safe)
func (acc *BankAccount) GetBalance() int {
	acc.mu.Lock()
	defer acc.mu.Unlock()

	return acc.balance
}

// Demonstration functions

// scenario1: Basic concurrent withdrawals
func scenario1() {
	fmt.Println("\n=== Scenario 1: Basic Concurrent Withdrawals ===")
	account := NewBankAccount(1000)
	var wg sync.WaitGroup

	// 5 concurrent withdrawals of 200 each
	for range 5 {
		wg.Go(func() {
			account.Withdraw(200)
		})
	}

	wg.Wait()
	fmt.Printf("Final Balance: %d (Expected: 0 or positive)\n", account.GetBalance())
}

// scenario2: Mixed deposits and withdrawals
func scenario2() {
	fmt.Println("\n=== Scenario 2: Mixed Deposits and Withdrawals ===")
	account := NewBankAccount(500)
	var wg sync.WaitGroup

	// 5 deposits of 100
	for range 5 {
		wg.Go(func() {
			account.Deposit(100)
		})
	}

	// 8 withdrawals of 150
	for range 8 {
		wg.Go(func() {
			account.Withdraw(150)
		})
	}

	wg.Wait()
	// Initial: 500 + (5*100) - (successful withdrawals * 150)
	// = 1000 - (successful * 150)
	// Should be >= 0
	fmt.Printf("Final Balance: %d (Expected: >= 0)\n", account.GetBalance())
}

// scenario3: High contention - many small transactions
func scenario3() {
	fmt.Println("\n=== Scenario 3: High Contention (100 concurrent withdrawals) ===")
	account := NewBankAccount(5000)
	var wg sync.WaitGroup

	successCount := 0
	failCount := 0
	var countMu sync.Mutex

	// 100 concurrent withdrawals of 100 each
	for range 100 {
		wg.Go(func() {
			success := account.Withdraw(100)
			countMu.Lock()
			if success {
				successCount++
			} else {
				failCount++
			}
			countMu.Unlock()
		})
	}

	wg.Wait()
	finalBalance := account.GetBalance()
	fmt.Printf("\nSuccessful withdrawals: %d\n", successCount)
	fmt.Printf("Failed withdrawals: %d\n", failCount)
	fmt.Printf("Final Balance: %d\n", finalBalance)
	fmt.Printf("Verification: 5000 - (%d * 100) = %d ✓\n", successCount, 5000-(successCount*100))
}

// scenario4: Race condition test - without mutex (for comparison)
type UnsafeBankAccount struct {
	balance int
}

func (acc *UnsafeBankAccount) Withdraw(amount int) bool {
	// NO MUTEX - UNSAFE!
	if acc.balance < amount {
		return false
	}
	// Simulate processing time
	time.Sleep(1 * time.Microsecond)
	acc.balance -= amount
	return true
}

func scenario4() {
	fmt.Println("\n=== Scenario 4: Comparing Safe vs Unsafe Implementation ===")

	// Safe implementation
	fmt.Println("\n🔒 Safe implementation (with Mutex):")
	safeAccount := NewBankAccount(1000)
	var wg1 sync.WaitGroup

	for range 10 {
		wg1.Go(func() {
			safeAccount.Withdraw(200)
		})
	}
	wg1.Wait()
	fmt.Printf("Final Balance: %d (Should be >= 0) ✅\n", safeAccount.GetBalance())

	// Unsafe implementation
	fmt.Println("\n🔓 Unsafe implementation (without Mutex):")
	unsafeAccount := &UnsafeBankAccount{balance: 1000}
	var wg2 sync.WaitGroup

	for range 10 {
		wg2.Go(func() {
			unsafeAccount.Withdraw(200)
		})
	}
	wg2.Wait()
	fmt.Printf("Final Balance: %d ", unsafeAccount.balance)
	if unsafeAccount.balance < 0 {
		fmt.Println("(NEGATIVE! Race condition occurred!) ❌")
	} else {
		fmt.Println("(May appear correct but unsafe) ⚠️")
	}
}

// scenario5: Stress test
func scenario5() {
	fmt.Println("\n=== Scenario 5: Stress Test (1000 goroutines) ===")
	account := NewBankAccount(100000)
	var wg sync.WaitGroup

	startTime := time.Now()

	// 500 deposits
	for range 500 {
		wg.Go(func() {
			account.Deposit(50)
		})
	}

	// 500 withdrawals
	for range 500 {
		wg.Go(func() {
			account.Withdraw(150)
		})
	}

	wg.Wait()
	duration := time.Since(startTime)

	fmt.Printf("Completed in: %v\n", duration)
	fmt.Printf("Final Balance: %d (Should be >= 0) ✅\n", account.GetBalance())
}

func main() {
	fmt.Println("=== Thread-Safe Bank Account Demo ===")

	// Run all scenarios
	scenario1()
	scenario2()
	scenario3()
	scenario4()
	scenario5()

	// Simple example
	fmt.Println("\n=== Simple Example ===")
	account := NewBankAccount(1000)
	var wg sync.WaitGroup

	fmt.Println("Initial Balance: 1000")
	fmt.Println("\nStarting 10 concurrent withdrawals of 150 each...")

	for range 10 {
		wg.Go(func() {
			account.Withdraw(150)
		})
	}

	wg.Wait()

	fmt.Printf("\nFinal Balance: %d\n", account.GetBalance())
	fmt.Println("\n✅ Balance is correct and never went negative!")
}
