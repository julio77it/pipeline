package pipeline_test

import (
	"sync"
	"testing"

	"github.com/julio77it/pipeline"

	"github.com/stretchr/testify/assert"
)

func inc(num int) int {
	num++

	return num
}

func giveBack(back *int, wg *sync.WaitGroup) func(num int) {
	return func(num int) {
		*back = num
		wg.Done()
	}
}

func TestRunner(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	var output int

	pipe, err := pipeline.Chain(
		inc,
		inc,
		inc,
		inc,
		inc,
		giveBack(&output, &wg),
	)
	assert.NoError(t, err)
	assert.NotNil(t, pipe)

	var input int = 5

	pipe.Process(input)
	wg.Wait()

	assert.Equal(t, input+5, output)
}
