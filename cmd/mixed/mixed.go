package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/julio77it/pipeline"
)

func sum(a, b int) int {
	fmt.Printf("Sum %d %d\n", a, b)
	return a + b
}

func toString(num int) string {
	fmt.Printf("ToString %d\n", num)
	return fmt.Sprintf("%d", num)
}

func toInt(str string) int {
	fmt.Printf("ToInt %s\n", str)
	n, _ := strconv.Atoi(str)
	return n
}

func double(num int) (int, int) {
	fmt.Printf("Double %d\n", num)
	return num, num
}

func giveBack(back *int) func(num int) {
	return func(num int) {
		*back = num
	}
}

func main() {
	var input, output int

	flag.IntVar(&input, "n", 0, "num to process")
	flag.Parse()

	pipe, err := pipeline.Chain(
		double, // input(int)       -> output(int,int)

		sum, // input(int,int)      -> output(int+int)

		double, // input(int)       -> output(int,int)

		double, // input(int)       -> output(int,int)
		// [second int from previous will be ignored]

		sum, // input(int,int)      -> output(int+int)

		toString, // input(int)     -> output(string)

		toInt, // input(string)     -> output(int)

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
