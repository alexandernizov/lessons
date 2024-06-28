package main

import (
	"alexandernizov/lessons/6/pool"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func main() {
	TaskWError()
}

func TaskWError() {
	tasksCount := 50
	tasks := make([]pool.Task, 0, tasksCount)

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
	maxErrorsCount := 30
	err := pool.Run(tasks, workersCount, maxErrorsCount)
	time.Sleep(1 * time.Second)
	fmt.Printf("ошибка: %v\n", err)
	fmt.Printf("количество ошибок: %v\n", maxErrorsCount)
	fmt.Printf("количество воркеров: %v\n", workersCount)
	fmt.Printf("количество запущенных задач: %v\n", runTasksCount)

}

func TaskWOError() {
	tasksCount := 50
	tasks := make([]pool.Task, 0, tasksCount)

	var runTasksCount int32
	var sumTime time.Duration

	for i := 0; i < tasksCount; i++ {
		taskSleep := time.Millisecond * time.Duration(rand.Intn(1))
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
	err := pool.Run(tasks, workersCount, maxErrorsCount)
	elapsedTime := time.Since(start)

	fmt.Printf("ошибка: %v\n", err)
	fmt.Printf("количество ошибок: %v\n", maxErrorsCount)
	fmt.Printf("количество воркеров: %v\n", workersCount)
	fmt.Printf("количество запущенных задач: %v\n", runTasksCount)
	fmt.Printf("всего времени сколько спят задачи: %v\n", sumTime)
	fmt.Printf("всего времени затрачено на выполнение задач: %v\n", elapsedTime)
}
