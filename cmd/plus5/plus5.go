package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/julio77it/pipeline"
)

func inc(num int) int {
	num++

	return num
}

func print(num int) int {
	fmt.Println("print ", num)

	return num
}

func giveBack(back *int) func(num int) {
	return func(num int) {
		*back = num
	}
}

func main() {
	var input, output int

	flag.IntVar(&input, "n", 0, "num to add 5")
	flag.Parse()

	pipe, err := pipeline.Chain(
		inc,
		inc,
		inc,
		inc,
		inc,
		print,
		giveBack(&output),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	pipe.Process(input)

	time.Sleep(time.Second)

	fmt.Println("Input : ", input)
	fmt.Println("Output : ", output)
}
