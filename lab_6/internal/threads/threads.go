package threads

import (
	"fmt"
	"sync"
)

type TurnstileCounter struct {
	count int
	mu    sync.Mutex
}

func (tc *TurnstileCounter) Increment() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.count++
}

func (tc *TurnstileCounter) Decrement() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.count--
}

func (tc *TurnstileCounter) GetCount() int {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	return tc.count
}

func ThreadsTest() {
	turnstileCounter := TurnstileCounter{}

	// Использование примитива синхронизации для увеличения и уменьшения счетчика
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			turnstileCounter.Increment()
		}()
	}
	wg.Wait()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			turnstileCounter.Decrement()
		}()
	}
	wg.Wait()

	fmt.Println("Счетчик людей на турникете:", turnstileCounter.GetCount())
}
