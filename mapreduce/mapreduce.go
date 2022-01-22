package main

import (
	"fmt"
	"go-zero-demo/mapreduce/goruntine"
	"log"

	"github.com/kevwan/mapreduce"
)

func main() {
	//mapReduceSample()
	//example.CountFunc()
	//example.CountFuncSample()
	//example.TestRange()
	//example.TestControlGoNum()
	//example.TestControlGoNum2()
	//goruntine.TestAtomic()
	//goruntine.TestAtomicAdd()
	//goruntine.TestSConfig()
	//goruntine.TestSConfigUseRWMutex()
	goruntine.TestSConfigUseAtomic()
}

func mapReduceSample() {
	val, err := mapreduce.MapReduce(func(source chan<- interface{}) {
		// generator
		for i := 0; i < 1000; i++ {
			source <- i
		}
	}, func(item interface{}, writer mapreduce.Writer, cancel func(error)) {
		// mapper
		i := item.(int)
		writer.Write(i * i)
		/*if i == 666 {
			cancel(errors.New("tet cancel"))
		}*/
	}, func(pipe <-chan interface{}, writer mapreduce.Writer, cancel func(error)) {
		// reducer
		var sum int
		for i := range pipe {
			sum += i.(int)
		}
		writer.Write(sum)
	}, mapreduce.WithWorkers(30))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mapReduceSample result:", val)
}
