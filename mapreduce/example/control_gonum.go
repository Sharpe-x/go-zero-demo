package example

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func TestControlGoNum() {
	jobsNum := 20
	var wg sync.WaitGroup
	var controlChan = make(chan int, 4)
	for i := 0; i < cap(controlChan); i++ {
		go func() {
			for controlEnum := range controlChan {
				fmt.Printf("hello %d\n", controlEnum)
				time.Sleep(time.Second)
				wg.Done()
			}
		}()
	}

	for i := 0; i < jobsNum; i++ {
		wg.Add(1)
		controlChan <- i
		fmt.Printf("index:%d,goruntine Num:%d\n", i, runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Printf("Done\n")
}

func TestControlGoNum2() {

	jobsNum := 50
	var wg sync.WaitGroup
	var controlChan = make(chan struct{}, 4)

	for i := 0; i < jobsNum; i++ {

		iTemp := i
		controlChan <- struct{}{}
		wg.Add(1)

		Go(func() {
			defer wg.Done()
			fmt.Printf("index:%d,goruntine Num:%d\n", iTemp, runtime.NumGoroutine())
			fmt.Printf("int = %d\n", iTemp)
			time.Sleep(time.Second)
			<-controlChan
		})
	}
	wg.Wait()
	fmt.Printf("Done\n")
}
