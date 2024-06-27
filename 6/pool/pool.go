package pool

import (
	"context"
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	//Каналы для задач и для результата
	tCh := make(chan Task, len(tasks))
	rCh := make(chan error, len(tasks))
	defer close(rCh)
	//Контекст для отмены выполнения воркеров, если мы получили слишком много ошибок
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Записали задачи
	for _, t := range tasks {
		tCh <- t
	}
	close(tCh)

	//WG для того, чтобы подождать, когда все воркеры завершат работу, иначе Run завершается до того, как выполняются воркеры
	//Run создавал канал для записи результатов, он же и закрывает. По-идее если сам воркер будет закрывать канал - тоже хорошо
	//Мы получим то, что RUN не будет дожидаться окончания работы всех воркеров, если ошибок > m. А оставшиеся горутины сами
	//себя завершат (сам пишет -> сам закрывает). И я где-то видел пример, где закрывали канал в Згорутинах, но не могу найти
	wg := sync.WaitGroup{}
	defer wg.Wait()

	//Запускаем воркеров
	for w := 0; w < n; w++ {
		wg.Add(1)
		go worker(ctx, tCh, rCh, &wg)
	}

	//Проверяем результат, если m == 0 - то не считаем ошибки
	errCount := 0
	for i := 0; i < len(tasks); i++ {
		err := <-rCh
		if err != nil {
			errCount++
		}
		if m > 0 && errCount >= m {
			cancel()
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}

func worker(ctx context.Context, tasks chan Task, results chan error, wg *sync.WaitGroup) {
	//По завершению своей работы воркер должен отчитаться, чтобы основная функция могла закрыть канал
	defer wg.Done()
	for {
		//Первый селект - если получаем отмену в контексте - то закрываемся
		select {
		case <-ctx.Done():
			return

		case t, ok := <-tasks:
			if !ok {
				return
			}
			err := t()
			//Второй селект, так как функция t() может выполняться долго, то к этому моменту уже канал мог быть закрыт
			select {
			case <-ctx.Done():
				return
			case results <- err:
			}
		}
	}
}
