package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int64

	sem := make(chan struct{}, int(pool))
	for i := 0; i < int(n); i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(item int64) {
			user := getOne(item)
			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			<-sem
			wg.Done()
		}(counter)
		counter++
	}
	wg.Wait()
	return res
}
