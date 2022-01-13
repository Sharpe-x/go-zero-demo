package example

import (
	"bufio"
	"fmt"
	"github.com/kevwan/mapreduce"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// GenerateFunc is used to let callers send elements into source.
//GenerateFunc func(source chan<- interface{})

var dirPath = "/Users/sharpe/workspace/dev_cloud/go_src/go-zero-demo/mapreduce"

func CountFuncSample() {
	val, err := mapreduce.MapReduce(func(source chan<- interface{}) {

		_ = filepath.Walk(dirPath, func(fileName string, f os.FileInfo, err error) error {

			if !f.IsDir() && path.Ext(fileName) == ".go" {
				source <- fileName
			}

			return nil
		})
	}, func(item interface{}, writer mapreduce.Writer, cancel func(error)) {

		fmt.Printf("process item = %v\n", item)
		lines := make(chan string)
		go func() {
			file, err := os.Open(item.(string))
			if err != nil {
				return
			}
			defer file.Close()

			reader := bufio.NewReader(file)
			for {
				line, readErr := reader.ReadString('\n')
				if readErr == io.EOF {
					break
				}

				if !strings.HasPrefix(line, "#") {
					lines <- line
				}
			}
			close(lines)
		}()

		var result int
		for line := range lines {
			if strings.HasPrefix(strings.TrimSpace(line), "func") {
				result++
			}
		}

		writer.Write(result)

	}, func(pipe <-chan interface{}, writer mapreduce.Writer, cancel func(error)) {

		var sum int
		for num := range pipe {
			sum += num.(int)
		}

		writer.Write(sum)
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(val)
}
