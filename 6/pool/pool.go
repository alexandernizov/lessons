package pool

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	//Каналы для задач и для подсчета ошибок
	taskCh := make(chan Task)
	errCh := make(chan bool)
	//Канал для ожидания общего результата
	resCh := make(chan bool)

	//Запускаем воркеров
	wg := sync.WaitGroup{}

	for w := 0; w < n; w++ {
		wg.Add(1)
		go worker(taskCh, errCh, &wg)
	}

	//Функция ожидания заврешения воркеров
	go func() {
		wg.Wait()
		//Закрыли канал, когда воркеры уже в него ничего не пишут
		close(errCh)
	}()

	//Горутина записывает в канал следующую задачу
	go func() {

		errCount := 0
		taskNum := 0

		//Если ошибки произошли в самых последних задачах, то мы всё равно должны понять, в результате нам отдавать Err или nil?
		f := func() {
			for err := range errCh {
				if err {
					errCount++
				}
			}
			if errCount >= m {
				resCh <- true
			}
			close(resCh)
		}

		defer f()
		defer close(taskCh)

		for {
			select {
			//Здесь не даём запустить следующую задачу, если мы получили слишком много ошибок (количество запусков должно быть n+m)
			case <-errCh:
				errCount++
				if errCount >= m {
					return
				}
			case taskCh <- tasks[taskNum]:
				taskNum++
				if taskNum == len(tasks) {
					return
				}
			}
		}

	}()

	//Ждем общий результат по отработанным задачам
	err := <-resCh
	if err {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func worker(tasks chan Task, errCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		err := task()
		if err != nil {
			errCh <- true
		}
	}
}
