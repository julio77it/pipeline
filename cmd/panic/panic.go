package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/julio77it/pipeline"
)

func toString(num int) string {
	fmt.Printf("ToString %d\n", num)
	return fmt.Sprintf("%d", num)
}

func toInt(str string) int {
	fmt.Printf("ToInt %s\n", str)
	n, _ := strconv.Atoi(str)
	return n
}

func main() {
	var input int

	flag.IntVar(&input, "n", 0, "num to process")
	flag.Parse()

	pipe, err := pipeline.Chain(
		toString, // input(int) -> output(string)
		toInt,    // input(string) -> output(int)
		toInt,    // input(string) -> output(int)
		// but input is a string ! not compatible, panic !!!
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	pipe.Process(input)

	time.Sleep(time.Second)

}
