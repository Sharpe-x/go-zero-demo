package goruntine

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func loadNewConfig() Config {
	return Config{
		NodeName: "西安",
		Addr:     "1.1.1.1",
		Count:    rand.Int31(),
	}
}

func TestAtomic() {
	var config atomic.Value
	config.Store(loadNewConfig())
	var cond = sync.NewCond(&sync.Mutex{})

	go func() {
		for {
			time.Sleep(time.Duration(5+rand.Int63n(5)) * time.Second)
			config.Store(loadNewConfig())
			cond.Broadcast()
		}
	}()

	go func() {
		for {
			cond.L.Lock()
			cond.Wait()
			c := config.Load().(Config)
			fmt.Printf("new config is :%v \n", c)
			cond.L.Unlock()
		}
	}()

	select {}
}

func TestAtomicAdd() {

	var i uint32
	newI := atomic.AddUint32(&i, 100)
	fmt.Printf("the new i is %d\n", newI)

	newI = atomic.AddUint32(&i, ^uint32(10-1))
	fmt.Printf("the new i is %d\n", newI)

	if atomic.CompareAndSwapUint32(&i, 91, 111) {
		fmt.Printf("the i is %d\n", i)
	}

	if atomic.CompareAndSwapUint32(&i, 90, 110) {
		fmt.Printf("the i is %d\n", i)
	}

	atomic.SwapUint32(&i, 22223)
	fmt.Printf("the i is %d\n", atomic.LoadUint32(&i))

	atomic.StoreUint32(&i, 44444)
	fmt.Printf("the i is %d\n", atomic.LoadUint32(&i))

	time.Sleep(10 * time.Second)

}
