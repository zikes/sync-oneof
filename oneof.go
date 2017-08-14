package oneof

import "sync"

func OneOf(f ...func() bool) bool {
	var wg sync.WaitGroup

	result := make(chan bool, 1)

	for _, fn := range f {
		wg.Add(1)
		go func(fn func() bool) {
			res := fn()
			if res {
				result <- true
			}
			wg.Done()
		}(fn)
	}

	go func() {
		wg.Wait()
		result <- false
	}()

	return <-result
}
