package mutex

import (
	"sync"
)

var (
	counter int = 0
	lock    sync.Mutex
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
}

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()
	counter++
	wg.Done()
}
