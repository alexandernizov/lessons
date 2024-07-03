package pipelineEasyWay

// Здесь реализация простого пайплайна, правила:
// 	1. Если с входящего канала больше нечего считать, то функция выполняет свою последнюю таску и закрывается самостоятельно
// 	2. Пока без done \ panic
//	3. Канал In закрывает тот, кто его создал, тем самым говорит, что больше этот пайплайн нам не нужен

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages {
		out = stage(out)
	}

	// Place your code here.
	return out
}

func StageIncrementInt(in <-chan interface{}) (out <-chan interface{}) {
	res := make(chan interface{})
	go func() {
		defer close(res)
		for value := range in {
			switch v := value.(type) {
			case int:
				v = v + 1
				res <- v
			default:
			}
		}
	}()
	return res
}

func StageSquareInt(in <-chan interface{}) (out <-chan interface{}) {
	res := make(chan interface{})
	go func() {
		defer close(res)
		for value := range in {
			switch v := value.(type) {
			case int:
				v = v * v
				res <- v
			default:
			}
		}
	}()
	return res
}
