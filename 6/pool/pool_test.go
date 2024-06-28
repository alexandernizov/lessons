package pool

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

func TestRunWError(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 3
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)
		fmt.Println(err)
		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})
}

func TestRunWoError(t *testing.T) {
	defer goleak.VerifyNone(t)
	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)
		require.Equal(t, int32(tasksCount), runTasksCount, "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

func TestWithoutError(t *testing.T) {
	//Тест на утечку горутин
	defer goleak.VerifyNone(t)

	jobsCount := 20

	jobs := []Task{}
	for i := 0; i < jobsCount; i++ {
		jobs = append(jobs, jobWithoutError)
	}

	err := Run(jobs, 1, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = Run(jobs, 8, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = Run(jobs, 1, 1)
	assert.Equal(t, nil, err, "expected 0 error")
	err = Run(jobs, 8, 1)
	assert.Equal(t, nil, err, "expected 0 error")
}

func TestWithAllError(t *testing.T) {
	//Тест на утечку горутин
	defer goleak.VerifyNone(t)

	jobsCount := 20

	jobs := []Task{}
	for i := 0; i < jobsCount; i++ {
		jobs = append(jobs, jobWithError)
	}

	err := Run(jobs, 1, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = Run(jobs, 8, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = Run(jobs, 1, 1)
	assert.Equal(t, ErrErrorsLimitExceeded, err, "expected error")
	err = Run(jobs, 8, 1)
	assert.Equal(t, ErrErrorsLimitExceeded, err, "expected error")
	err = Run(jobs, 1, 2)
	assert.Equal(t, ErrErrorsLimitExceeded, err, "expected error")
	err = Run(jobs, 8, 2)
	assert.Equal(t, ErrErrorsLimitExceeded, err, "expected error")
}

func TestWithPartuallyError(t *testing.T) {
	//Тест на утечку горутин
	defer goleak.VerifyNone(t)

	jobsCount := 20

	jobs := []Task{}
	for i := 0; i < (jobsCount / 2); i++ {
		jobs = append(jobs, jobWithError)
	}
	for i := 0; i < (jobsCount / 2); i++ {
		jobs = append(jobs, jobWithoutError)
	}

	err := Run(jobs, 1, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = Run(jobs, 8, 0)
	assert.Equal(t, nil, err, "expected 0 error")
	err = Run(jobs, 1, 1)
	assert.Equal(t, ErrErrorsLimitExceeded, err, "expected error")
	err = Run(jobs, 8, 1)
	assert.Equal(t, ErrErrorsLimitExceeded, err, "expected error")
	err = Run(jobs, 1, 2)
	assert.Equal(t, ErrErrorsLimitExceeded, err, "expected error")
	err = Run(jobs, 8, 2)
	assert.Equal(t, ErrErrorsLimitExceeded, err, "expected error")
}

func jobWithError() error {
	i := 1_000_000_000
	rnd := int64(rand.Intn(i))

	//some difficult work
	for i := 0; i < int(rnd); i++ {
		_ = i
	}

	return errors.New("error")
}

func jobWithoutError() error {
	i := 1_000_000_000
	rnd := int64(rand.Intn(i))

	//some difficult work
	for i := 0; i < int(rnd); i++ {
		_ = i
	}

	return nil
}
