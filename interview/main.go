package main

import (
	"context"
	"fmt"
)

// Task 任务
type Task struct {
	Name string
}

// Do 任务具体执行方法
func (t *Task) Do(ctx context.Context, i int) error {
	fmt.Printf("Task Name is：%s，Number is: %d \n", t.Name, i)
	return nil
}

// ConcurrencyDoTask 并发执行任务
func ConcurrencyDoTask(ctx context.Context, tasks []*Task, concurrent int64) error {
	// todo 任务一次只能完成10个,遇到错误退出整个任务
	// 当不满十个的时候，需要添加任务执行

	errChan := make(chan error)
	// 单次任务执行最大的并发数
	countChan := make(chan int64, concurrent)

	ctx1, cancel := context.WithCancel(ctx)

	for i, task := range tasks {
		go func(ctx context.Context, task *Task, i int) {
			select {
			case <-ctx.Done():
				fmt.Println("退出")
				return
			default:
				// 执行任务的时候 往里面添加1
				countChan <- 1
				err := task.Do(ctx, i)
				if err != nil {
					errChan <- err
					return
				}
				<-countChan
			}
		}(ctx1, task, i)
	}
	for {
		select {
		case err := <-errChan:
			cancel()
			return err
		default:
			return nil
		}
	}

}

func main() {
	tasks := make([]*Task, 0)
	t1 := &Task{Name: "name1"}
	tasks = append(tasks, t1)
	t2 := &Task{Name: "name2"}
	tasks = append(tasks, t2)
	t3 := &Task{Name: "name3"}
	tasks = append(tasks, t3)
	t4 := &Task{Name: "name4"}
	tasks = append(tasks, t4)
	t5 := &Task{Name: "name5"}
	tasks = append(tasks, t5)
	t6 := &Task{Name: "name6"}
	tasks = append(tasks, t6)
	t7 := &Task{Name: "name7"}
	tasks = append(tasks, t7)
	t8 := &Task{Name: "name8"}
	tasks = append(tasks, t8)
	t9 := &Task{Name: "name9"}
	tasks = append(tasks, t9)
	t10 := &Task{Name: "name10"}
	tasks = append(tasks, t10)
	t11 := &Task{Name: "name11"}
	tasks = append(tasks, t11)
	t12 := &Task{Name: "name12"}
	tasks = append(tasks, t12)
	t13 := &Task{Name: "name13"}
	tasks = append(tasks, t13)
	t14 := &Task{Name: "name14"}
	tasks = append(tasks, t14)
	t15 := &Task{Name: "name15"}
	tasks = append(tasks, t15)
	ctx := context.Background()
	err := ConcurrencyDoTask(ctx, tasks, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
