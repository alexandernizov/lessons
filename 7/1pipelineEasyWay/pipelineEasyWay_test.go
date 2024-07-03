package pipelineEasyWay_test

import (
	"sync/atomic"
	"testing"
	"time"

	pipeline "github.com/alexandernizov/lessons/7/1pipelineEasyWay"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestExecutePipeline(t *testing.T) {
	defer goleak.VerifyNone(t)

	in := make(chan interface{})
	done := make(<-chan interface{})

	inputsCount := 10
	go func() {
		defer close(in)
		for i := 0; i < inputsCount; i++ {
			in <- i
		}
	}()
	results := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}

	stages := []pipeline.Stage{}
	stages = append(stages, pipeline.StageIncrementInt)
	stages = append(stages, pipeline.StageSquareInt)

	out := pipeline.ExecutePipeline(in, done, stages...)
	for _, v := range results {
		assert.Equal(t, v, <-out, "we got incorrect value")
	}
}

func TestConcurencyOfPipeline(t *testing.T) {
	in := make(chan interface{})
	done := make(<-chan interface{})

	inputsCount := 10
	go func() {
		defer close(in)
		for i := 0; i < inputsCount; i++ {
			in <- i
		}
	}()

	var runTasksCount int32
	var sumTime time.Duration
	testFn := func(in <-chan interface{}) (out <-chan interface{}) {
		res := make(chan interface{})
		go func() {
			defer close(res)
			for v := range in {
				taskSleep := time.Millisecond * 100 //time.Duration(rand.Intn(100))
				sumTime += taskSleep
				atomic.AddInt32(&runTasksCount, 1)
				time.Sleep(taskSleep)
				res <- v
			}
		}()
		return res
	}

	stages := []pipeline.Stage{}
	stages = append(stages, testFn)
	stages = append(stages, testFn)
	stages = append(stages, testFn)
	stages = append(stages, testFn)
	stages = append(stages, testFn)

	start := time.Now()
	out := pipeline.ExecutePipeline(in, done, stages...)
	for range out {
	}
	elapsedTime := time.Since(start)
	require.Equal(t, int32(len(stages)*inputsCount), runTasksCount, "not all tasks were completed")
	require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
}
