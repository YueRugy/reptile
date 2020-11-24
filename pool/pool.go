package pool

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock            sync.Mutex
	defaultPoolSize = 10
)

type Product struct {
	Jobs chan *Data
}

type Consumer struct {
	count   int
	WorkNum chan struct{}
	Handler func(jobs *Data)
}

type Task struct {
	Consumer
	Product
}

//NewTask
func NewTask(workNum int, handler func(jobs *Data)) *Task {

	task := &Task{
		Consumer: Consumer{
			count:   defaultPoolSize,
			WorkNum: make(chan struct{}, workNum),
			Handler: handler,
		},
		Product: Product{
			Jobs: make(chan *Data, workNum*10),
		},
	}

	for i := 0; i < defaultPoolSize; i++ {
		task.WorkNum <- struct{}{}
	}

	return task
}

func (task *Task) DisposeData() {

	for true {
		select {
		case job:=<-task.Product.Jobs:
			go func() {
				if len(task.WorkNum) == 0 && task.count < cap(task.WorkNum) {
					lock.Lock()
					task.count++
					task.WorkNum <- struct{}{}
					lock.Unlock()
				}
				<-task.WorkNum
				task.Handler(job)
				defer func() {
					//fmt.Println("count"+strconv.Itoa(task.count))
					task.WorkNum <- struct{}{}
				}()
			}()
		case <- time.After(time.Second*5):
			fmt.Println("dont get job sleep 5 second")
		}

	}
}

type Data struct {
	Url string
	Title string
}

func (p *Product) AddData(url ,title string) {
	data := &Data{
		Url: url,
		Title:title,
	}
	p.Jobs <- data
}
