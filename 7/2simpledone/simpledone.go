package simpledone

import "fmt"

//Реализация простого Done
// Идея: так как стейдж работает до того момента, пока работает инпут - то:
// 	- создаем свой инпут канал и начитываем в него значения из исходного инпута
// 	- добавляем селект, в котором происходит событие: либо из внешенго инпута кладется значение в свой собственный инпут, либо приходит done и все закрывается
//  - по нисходящей закрываются каналы стейджей
// 	- те значения, которые уже лежат в стейджах доработают до конца
// 	- последний канал последнего стейджа закроется
// 	- и все будет ок

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	//Создадим свой канал input, который будем закрывать когда получили done или значения в исходном канале закончились

	newInput := make(chan interface{})
	go func() {
		defer close(newInput)
		for {
			select {
			case val, ok := <-in:
				if !ok {
					fmt.Println("внешний input закрылся")
					return
				}
				newInput <- val
			case <-done:
				fmt.Println("считали done")
				return
			}
		}
	}()

	out := make(<-chan interface{})
	out = newInput
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
