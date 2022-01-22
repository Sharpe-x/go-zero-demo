package goruntine

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SConfig struct {
	data []int
}

func CreateSConfig() {
	cfg := &SConfig{}

	go func() {
		i := 0
		for {
			i++
			cfg.data = []int{i, i + 1, i + 2, i + 3, i + 4}
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// 这里打印出来的数组不连续 因为产生了data race
				fmt.Printf("Sconfig data is %#v\n", cfg)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func CreateSConfigUseRWMutex() {
	cfg := &SConfig{}
	var m sync.RWMutex

	go func() {
		i := 0
		for {
			i++
			m.Lock()
			cfg.data = []int{i, i + 1, i + 2, i + 3, i + 4}
			m.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// 这里打印出来的数组连续 因为使用了读写锁
				m.RLock()
				fmt.Printf("Sconfig data is %#v\n", cfg)
				m.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func CreateSConfigUseAtomic() {
	//cfg := &SConfig{}
	var v atomic.Value

	go func() {
		i := 0
		for {
			i++
			cfg := &SConfig{
				data: []int{i, i + 1, i + 2, i + 3, i + 4},
			}
			v.Store(cfg)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// 这里打印出来的数组连续 因为使用了读写锁
				// v.load 有可能还没有数据
				fmt.Printf("Sconfig data is %#v\n", v.Load())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
